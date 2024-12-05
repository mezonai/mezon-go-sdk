package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"mezon-sdk/mezon-protobuf/mezon/v2/common/rtapi"
	"sync"

	"time"

	"mezon-sdk/constants"
	"mezon-sdk/utils"

	"github.com/asticode/go-astiav"
	"github.com/pion/webrtc/v4"
)

var (
	mapChannelRtcconnection sync.Map // map[channelId]*RTCConnection
)

type RTCConnection struct {
	peer *webrtc.PeerConnection
	ws   *WSConnection

	inputFormatContext *astiav.FormatContext

	decodeCodecContext *astiav.CodecContext
	decodePacket       *astiav.Packet
	decodeFrame        *astiav.Frame
	videoStream        *astiav.Stream

	softwareScaleContext *astiav.SoftwareScaleContext
	scaledFrame          *astiav.Frame
	encodeCodecContext   *astiav.CodecContext
	encodePacket         *astiav.Packet

	pts int64
}

func NewRTCConnection(config webrtc.Configuration, channelId, clanId string) error {
	peerConnection, err := webrtc.NewPeerConnection(config)
	if err != nil {
		return err
	}
	peerConnection.OnICEConnectionStateChange(func(connectionState webrtc.ICEConnectionState) {
		fmt.Printf("Connection State has changed %s \n", connectionState.String())
	})

	// Create a video track
	videoTrack, err := webrtc.NewTrackLocalStaticRTP(webrtc.RTPCodecCapability{MimeType: "video/h264"}, "video", "pion")
	if err != nil {
		return err
	}
	// Create a audio track
	audioTrack, err := webrtc.NewTrackLocalStaticRTP(webrtc.RTPCodecCapability{MimeType: "video/h264"}, "video", "pion")
	if err != nil {
		return err
	}
	_, err = peerConnection.AddTrack(audioTrack)
	if err != nil {
		return err
	}

	mapChannelRtcconnection.Store(channelId, &RTCConnection{
		peer: peerConnection,
	})

	// TODO: connect websocket

	// Start pushing buffers on these tracks
	if rtcc, ok := mapChannelRtcconnection.Load(channelId); ok {
		rtcc.(*RTCConnection).writeH264ToTrack(videoTrack)
	}

	// Block forever
	select {}
}

func (c *RTCConnection) onWebsocketEvent(event *rtapi.Envelope) error {
	switch event.Message.(type) {
	case *rtapi.Envelope_JoinPttChannel:
		eventData := event.GetJoinPttChannel()
		switch eventData.DataType {
		case constants.WEBRTC_SDP_ANSWER:
			unzipData, _ := utils.GzipUncompress(eventData.JsonData)
			var answer webrtc.SessionDescription
			err := json.Unmarshal([]byte(unzipData), &answer)
			if err != nil {
				return err
			}

			return c.recvAnswer(answer)

		case constants.WEBRTC_ICE_CANDIDATE:

			var i webrtc.ICECandidateInit
			err := json.Unmarshal([]byte(eventData.JsonData), &i)
			if err != nil {
				return err
			}

			return c.addICECandidate(i)
		}

	}
	return nil
}

func (c *RTCConnection) sendOffer() error {
	offer, err := c.peer.CreateOffer(nil)
	if err != nil {
		return err
	}
	if err := c.peer.SetLocalDescription(offer); err != nil {
		return err
	}

	byteJson, _ := json.Marshal(offer)
	dataEnc, _ := utils.GzipCompress(string(byteJson))

	//TODO: send ws , gzip compress data
	return c.ws.SendMessage(&rtapi.Envelope{
		Cid: "",
		Message: &rtapi.Envelope_JoinPttChannel{
			JoinPttChannel: &rtapi.JoinPTTChannel{
				JsonData: dataEnc,
				DataType: constants.WEBRTC_SDP_OFFER,
			},
		},
	})
}

func (c *RTCConnection) recvAnswer(answer webrtc.SessionDescription) error {
	if err := c.peer.SetRemoteDescription(answer); err != nil {
		return err
	}
	return nil
}

func (c *RTCConnection) onICECandidate(i *webrtc.ICECandidate, channelId, userId string) error {
	if i == nil {
		return nil
	}
	// If you are serializing a candidate make sure to use ToJSON
	// Using Marshal will result in errors around `sdpMid`
	candidateString, err := json.Marshal(i.ToJSON())
	if err != nil {
		return err
	}

	return c.ws.SendMessage(&rtapi.Envelope{Message: &rtapi.Envelope_JoinPttChannel{JoinPttChannel: &rtapi.JoinPTTChannel{
		ReceiverId: userId,
		DataType:   constants.WEBRTC_ICE_CANDIDATE,
		JsonData:   string(candidateString),
		ChannelId:  channelId,
	}}})
}

func (c *RTCConnection) addICECandidate(i webrtc.ICECandidateInit) error {
	return c.peer.AddICECandidate(i)
}

func (c *RTCConnection) writeH264ToTrack(track *webrtc.TrackLocalStaticRTP) error {
	astiav.RegisterAllDevices()

	c.initTestSrc()
	defer c.freeVideoCoding()

	h264FrameDuration := time.Duration(float64(time.Second) / c.videoStream.AvgFrameRate().Float64())

	ticker := time.NewTicker(h264FrameDuration)
	for ; true; <-ticker.C {
		c.decodePacket.Unref()

		// Read frame from lavfi
		if err := c.inputFormatContext.ReadFrame(c.decodePacket); err != nil {
			if errors.Is(err, astiav.ErrEof) {
				break
			}
			return err
		}

		c.decodePacket.RescaleTs(c.videoStream.TimeBase(), c.decodeCodecContext.TimeBase())

		// Send the packet
		if err := c.decodeCodecContext.SendPacket(c.decodePacket); err != nil {
			return err
		}

		for {
			// Read Decoded Frame
			if err := c.decodeCodecContext.ReceiveFrame(c.decodeFrame); err != nil {
				if errors.Is(err, astiav.ErrEof) || errors.Is(err, astiav.ErrEagain) {
					break
				}
				return err
			}

			// Init the Scaling+Encoding. Can't be started until we know info on input video
			c.initVideoEncoding()

			// Scale the video
			if err := c.softwareScaleContext.ScaleFrame(c.decodeFrame, c.scaledFrame); err != nil {
				return err
			}

			// We don't care about the PTS, but encoder complains if unset
			c.pts++
			c.scaledFrame.SetPts(c.pts)

			// Encode the frame
			if err := c.encodeCodecContext.SendFrame(c.scaledFrame); err != nil {
				return err
			}

			for {
				// Read encoded packets and write to file
				c.encodePacket = astiav.AllocPacket()
				if err := c.encodeCodecContext.ReceivePacket(c.encodePacket); err != nil {
					if errors.Is(err, astiav.ErrEof) || errors.Is(err, astiav.ErrEagain) {
						break
					}
					return err
				}

				// Write H264 to track
				if _, err := track.Write(c.encodePacket.Data()); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (c *RTCConnection) initTestSrc() error {
	if c.inputFormatContext = astiav.AllocFormatContext(); c.inputFormatContext == nil {
		return errors.New("Failed to AllocCodecContext")
	}

	// Open input
	if err := c.inputFormatContext.OpenInput("testsrc=size=640x480:rate=30", astiav.FindInputFormat("lavfi"), nil); err != nil {
		return err
	}

	// Find stream info
	if err := c.inputFormatContext.FindStreamInfo(nil); err != nil {
		return err
	}

	c.videoStream = c.inputFormatContext.Streams()[0]

	decodeCodec := astiav.FindDecoder(c.videoStream.CodecParameters().CodecID())
	if decodeCodec == nil {
		return errors.New("FindDecoder returned nil")
	}

	if c.decodeCodecContext = astiav.AllocCodecContext(decodeCodec); c.decodeCodecContext == nil {
		return errors.New("AllocCodecContext returned nil")
	}

	if err := c.videoStream.CodecParameters().ToCodecContext(c.decodeCodecContext); err != nil {
		return err
	}

	c.decodeCodecContext.SetFramerate(c.inputFormatContext.GuessFrameRate(c.videoStream, nil))

	if err := c.decodeCodecContext.Open(decodeCodec, nil); err != nil {
		return err
	}

	c.decodePacket = astiav.AllocPacket()
	c.decodeFrame = astiav.AllocFrame()
	return nil
}

func (c *RTCConnection) initVideoEncoding() (err error) {
	if c.encodeCodecContext != nil {
		return nil
	}

	h264Encoder := astiav.FindEncoder(astiav.CodecIDH264)
	if h264Encoder == nil {
		return errors.New("No H264 Encoder Found")
	}

	if c.encodeCodecContext = astiav.AllocCodecContext(h264Encoder); c.encodeCodecContext == nil {
		return errors.New("Failed to AllocCodecContext Decoder")
	}

	c.encodeCodecContext.SetPixelFormat(astiav.PixelFormatYuv420P)
	c.encodeCodecContext.SetSampleAspectRatio(c.decodeCodecContext.SampleAspectRatio())
	c.encodeCodecContext.SetTimeBase(astiav.NewRational(1, 30))
	c.encodeCodecContext.SetWidth(c.decodeCodecContext.Width())
	c.encodeCodecContext.SetHeight(c.decodeCodecContext.Height())

	encodeCodecContextDictionary := astiav.NewDictionary()
	if err := encodeCodecContextDictionary.Set("bf", "0", astiav.NewDictionaryFlags()); err != nil {
		return err
	}

	if err = c.encodeCodecContext.Open(h264Encoder, encodeCodecContextDictionary); err != nil {
		return err
	}

	c.softwareScaleContext, err = astiav.CreateSoftwareScaleContext(
		c.decodeCodecContext.Width(),
		c.decodeCodecContext.Height(),
		c.decodeCodecContext.PixelFormat(),
		c.decodeCodecContext.Width(),
		c.decodeCodecContext.Height(),
		astiav.PixelFormatYuv420P,
		astiav.NewSoftwareScaleContextFlags(astiav.SoftwareScaleContextFlagBilinear),
	)
	if err != nil {
		return err
	}

	c.scaledFrame = astiav.AllocFrame()
	return nil
}

func (c *RTCConnection) freeVideoCoding() {
	c.inputFormatContext.CloseInput()
	c.inputFormatContext.Free()

	c.decodeCodecContext.Free()
	c.decodePacket.Free()
	c.decodeFrame.Free()

	c.scaledFrame.Free()
	c.softwareScaleContext.Free()
	c.encodeCodecContext.Free()
	c.encodePacket.Free()
}
