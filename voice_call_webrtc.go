package mezonsdk

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"image/jpeg"
	"log"
	"sync"
	"time"

	"github.com/nccasia/mezon-go-sdk/constants"
	"github.com/nccasia/mezon-go-sdk/mezon-protobuf/mezon/v2/common/rtapi"
	"github.com/nccasia/mezon-go-sdk/utils"

	"github.com/pion/rtcp"
	"github.com/pion/rtp"
	"github.com/pion/rtp/codecs"
	"github.com/pion/webrtc/v4"
	"github.com/pion/webrtc/v4/pkg/media/samplebuilder"
	"golang.org/x/image/vp8"
)

var mapCallRtcConn sync.Map // map[channelId]*RTCConnection

type callRTCConn struct {
	peer       *webrtc.PeerConnection
	channelId  string
	receiverId string
	callerId   string

	audioTrack     *webrtc.TrackLocalStaticRTP
	rtpChan        chan *rtp.Packet
	snapShootCount int
}

type callService struct {
	botId          string
	ws             IWSConnection
	config         webrtc.Configuration
	snapShootCount int
	onImage        func(imgBase64 string) error
}

type ICallService interface {
	SetOnImage(onImage func(imgBase64 string) error, snapShootCount int)
	SendFile(channelId, filePath string) error
	OnWebsocketEvent(event *rtapi.Envelope) error
	GetRTCConnectionState(channelId string) webrtc.PeerConnectionState
}

func NewCallService(botId string, wsConn IWSConnection, config webrtc.Configuration) ICallService {
	return &callService{
		botId:  botId,
		ws:     wsConn,
		config: config,
	}
}

func (c *callService) newCallRTCConnection(channelId, receiverId string) (*callRTCConn, error) {

	peerConnection, err := webrtc.NewPeerConnection(c.config)
	if err != nil {
		return nil, err
	}

	// Create a audio track
	// audioTrack, err := webrtc.NewTrackLocalStaticRTP(webrtc.RTPCodecCapability{MimeType: webrtc.MimeTypeOpus}, fmt.Sprintf("audio_opus_%s", channelId), fmt.Sprintf("audio_opus_%s", channelId))
	// if err != nil {
	// 	return nil, err
	// }
	// _, err = peerConnection.AddTrack(audioTrack)
	// if err != nil {
	// 	return nil, err
	// }

	// save to store
	rtcConnection := &callRTCConn{
		peer:       peerConnection,
		channelId:  channelId,
		receiverId: receiverId,
		callerId:   c.botId,
		// audioTrack:     audioTrack,
		snapShootCount: c.snapShootCount,
		rtpChan:        make(chan *rtp.Packet),
	}
	mapCallRtcConn.Store(channelId, rtcConnection)

	peerConnection.OnICEConnectionStateChange(func(state webrtc.ICEConnectionState) {
		c.onICEConnectionStateChange(state, channelId, receiverId)
	})
	peerConnection.OnICECandidate(func(i *webrtc.ICECandidate) {
		c.onICECandidate(i, channelId, c.botId, receiverId)
	})

	peerConnection.OnTrack(func(track *webrtc.TrackRemote, receiver *webrtc.RTPReceiver) {

		// save image by time receive track
		if rtcConnection.snapShootCount > 0 && track.Kind() == webrtc.RTPCodecTypeVideo {
			// Send a PLI on an interval so that the publisher is pushing a keyframe every rtcpPLIInterval
			go func() {
				ticker := time.NewTicker(time.Second * 3)
				for range ticker.C {
					errSend := peerConnection.WriteRTCP([]rtcp.Packet{&rtcp.PictureLossIndication{MediaSSRC: uint32(track.SSRC())}})
					if errSend != nil {
						return
					}
				}
			}()

			for {
				// Read RTP Packets in a loop
				rtpPacket, _, readErr := track.ReadRTP()
				if readErr != nil {
					log.Printf("track read rtp error: %+v \n", readErr)
					return
				}

				// Use a lossy channel to send packets to snapshot handler
				// We don't want to block and queue up old data
				select {
				case rtcConnection.rtpChan <- rtpPacket:
				default:
				}
			}
		}
	})

	return rtcConnection, nil
}

func (c *callService) OnWebsocketEvent(event *rtapi.Envelope) error {
	switch event.Message.(type) {
	case *rtapi.Envelope_WebrtcSignalingFwd:
		eventData := event.GetWebrtcSignalingFwd()

		if eventData.ReceiverId != c.botId {
			return nil
		}

		var rtcConn *callRTCConn
		rtcConnAny, ok := mapCallRtcConn.Load(eventData.ChannelId)
		if ok {
			rtcConn = rtcConnAny.(*callRTCConn)
		}

		switch eventData.DataType {
		case constants.WEBRTC_SDP_OFFER:
			if !ok {
				var err error
				rtcConn, err = c.newCallRTCConnection(eventData.ChannelId, eventData.CallerId)
				if err != nil {
					return err
				}
			}

			unzipData, _ := utils.GzipUncompress(eventData.JsonData)
			var offer webrtc.SessionDescription
			err := json.Unmarshal([]byte(unzipData), &offer)
			if err != nil {
				return err
			}
			if err := rtcConn.peer.SetRemoteDescription(offer); err != nil {
				return err
			}

			answer, err := rtcConn.peer.CreateAnswer(nil)
			if err != nil {
				return err
			}

			if err := rtcConn.peer.SetLocalDescription(answer); err != nil {
				return err
			}
			// TODO: ws send answer
			answerBytes, _ := json.Marshal(answer)
			zipData, _ := utils.GzipCompress(string(answerBytes))
			c.ws.SendMessage(&rtapi.Envelope{Message: &rtapi.Envelope_WebrtcSignalingFwd{WebrtcSignalingFwd: &rtapi.WebrtcSignalingFwd{
				DataType:   constants.WEBRTC_SDP_ANSWER,
				JsonData:   zipData,
				ChannelId:  eventData.ChannelId,
				CallerId:   rtcConn.callerId,
				ReceiverId: rtcConn.receiverId,
			}}})

		case constants.WEBRTC_ICE_CANDIDATE:

			var i webrtc.ICECandidateInit
			err := json.Unmarshal([]byte(eventData.JsonData), &i)
			if err != nil {
				return err
			}

			if ok {
				rtcConn.peer.AddICECandidate(i)
			}
		}

	}
	return nil
}

func (c *callService) GetRTCConnectionState(channelId string) webrtc.PeerConnectionState {
	rtcConn, ok := mapCallRtcConn.Load(channelId)
	if !ok {
		return webrtc.PeerConnectionStateUnknown
	}

	return rtcConn.(*callRTCConn).peer.ConnectionState()
}

func (c *callService) onICECandidate(i *webrtc.ICECandidate, channelId, callerId, receiverId string) error {
	if i == nil {
		return nil
	}
	// If you are serializing a candidate make sure to use ToJSON
	// Using Marshal will result in errors around `sdpMid`
	candidateString, err := json.Marshal(i.ToJSON())
	if err != nil {
		return err
	}

	return c.ws.SendMessage(&rtapi.Envelope{Message: &rtapi.Envelope_WebrtcSignalingFwd{WebrtcSignalingFwd: &rtapi.WebrtcSignalingFwd{
		DataType:   constants.WEBRTC_ICE_CANDIDATE,
		JsonData:   string(candidateString),
		ChannelId:  channelId,
		CallerId:   callerId,
		ReceiverId: receiverId,
	}}})
}

func (c *callService) onICEConnectionStateChange(state webrtc.ICEConnectionState, channelId, receiverId string) {
	log.Printf("Connection State has changed %s \n", state.String())

	rtcConn, ok := mapCallRtcConn.Load(channelId)
	if !ok {
		return
	}

	if rtcConn.(*callRTCConn).peer == nil {
		return
	}

	switch state {
	case webrtc.ICEConnectionStateConnected:
		rtcConn.(*callRTCConn).saveTrackToImage(c.onImage, receiverId)

	case webrtc.ICEConnectionStateClosed:

		if rtcConn.(*callRTCConn).peer.ConnectionState() != webrtc.PeerConnectionStateClosed {
			rtcConn.(*callRTCConn).peer.Close()
		}

		mapCallRtcConn.Delete(channelId)
	}
}

func (c *callService) SendFile(channelId, filePath string) error {
	// Start pushing buffers on these tracks
	audioSrc := fmt.Sprintf("uridecodebin uri=%s ! audioconvert", filePath)

	rtcConn, ok := mapCallRtcConn.Load(channelId)
	if !ok {
		return errors.New("can not found call rtc connection")
	}

	rtcConn.(*callRTCConn).pipelineForCodec("mp3", []*webrtc.TrackLocalStaticRTP{rtcConn.(*callRTCConn).audioTrack}, audioSrc)

	return nil
}

func (c *callService) SetOnImage(onImage func(imgBase64 string) error, snapShootCount int) {
	c.snapShootCount = snapShootCount
	c.onImage = onImage
}

func (c *callRTCConn) pipelineForCodec(codecName string, tracks []*webrtc.TrackLocalStaticRTP, pipelineSrc string) {
	log.Printf("[pipelineForCodec] codecName: %s, len(track): %d, pipelineSrc: %s \n", codecName, len(tracks), pipelineSrc)

	//TODO: ffmpeg or convert file in bot repo
}

func (c *callRTCConn) saveTrackToImage(onImage func(imgBase64 string) error, receiverId string) error {
	log.Printf("[saveTrackToImage] receiverId: %s \n", receiverId)

	imageCount := 1
	sampleBuilder := samplebuilder.New(20, &codecs.VP8Packet{}, 90000)
	decoder := vp8.NewDecoder()

	for {

		// Pull RTP Packet from rtpChan
		sampleBuilder.Push(<-c.rtpChan)

		// Use SampleBuilder to generate full picture from many RTP Packets
		sample := sampleBuilder.Pop()
		if sample == nil {
			continue
		}

		// Read VP8 header.
		videoKeyframe := (sample.Data[0]&0x1 == 0)
		if !videoKeyframe {
			continue
		}

		// Begin VP8-to-image decode: Init->DecodeFrameHeader->DecodeFrame
		decoder.Init(bytes.NewReader(sample.Data), len(sample.Data))

		// Decode header
		if _, err := decoder.DecodeFrameHeader(); err != nil {
			log.Printf("decoder.DecodeFrameHeader error: %+v \n", err)
			return err
		}

		// Decode Frame
		img, err := decoder.DecodeFrame()
		if err != nil {
			log.Printf("decoder.DecodeFrame error: %+v \n", err)
			return err
		}

		// Log the image size to check resolution
		// log.Printf("Decoded image size: width=%d, height=%d", img.Bounds().Dx(), img.Bounds().Dy())

		// Encode to (RGB) jpeg
		buffer := new(bytes.Buffer)
		if err = jpeg.Encode(buffer, img, nil); err != nil {
			log.Printf("jpeg.Encode error: %+v \n", err)
			continue
		}

		// // Convert JPEG bytes to Base64
		base64Image := base64.StdEncoding.EncodeToString(buffer.Bytes())
		imageCount++

		// log.Printf("Generated Base64 for image %s", base64Image)
		if err := onImage(base64Image); err != nil {
			return err
		}

		sampleBuilder = samplebuilder.New(20, &codecs.VP8Packet{}, 90000)
		decoder = vp8.NewDecoder()

		if imageCount > c.snapShootCount {
			close(c.rtpChan)
			c.peer.Close()

			return nil
		}
	}
}
