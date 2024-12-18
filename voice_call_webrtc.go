package mezonsdk

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"image/jpeg"
	"io"
	"log"
	"os"
	"sync"
	"time"

	"github.com/nccasia/mezon-go-sdk/constants"
	"github.com/nccasia/mezon-go-sdk/mezon-protobuf/mezon/v2/common/rtapi"
	"github.com/nccasia/mezon-go-sdk/utils"
	"github.com/pion/webrtc/v4/pkg/media"

	"github.com/pion/rtcp"
	"github.com/pion/rtp"
	"github.com/pion/rtp/codecs"
	"github.com/pion/webrtc/v4"
	"github.com/pion/webrtc/v4/pkg/media/oggreader"
	"github.com/pion/webrtc/v4/pkg/media/samplebuilder"
	"golang.org/x/image/vp8"
)

var mapCallRtcConn sync.Map // map[channelId]*RTCConnection

type callRTCConn struct {
	peer       *webrtc.PeerConnection
	channelId  string
	receiverId string
	callerId   string

	checkinSuccessAudioFile string
	checkinFailAudioFile    string
	acceptCallAudioFile     string
	exitCallAudioFile       string
	audioTrack              *webrtc.TrackLocalStaticSample
	rtpChan                 chan *rtp.Packet
	snapShootCount          int
	isVideoCall             bool
}

type callService struct {
	botId                   string
	ws                      IWSConnection
	config                  webrtc.Configuration
	snapShootCount          int
	checkinSuccessAudioFile string
	checkinFailAudioFile    string
	acceptCallAudioFile     string
	exitCallAudioFile       string
	onImage                 func(imgBase64 string) error
}

type ICallService interface {
	SetOnImage(onImage func(imgBase64 string) error, snapShootCount int)
	SetCheckinSuccessFileAudio(filePath string)
	SetCheckinFailFileAudio(filePath string)
	SetAcceptCallFileAudio(filePath string)
	SetExitCallFileAudio(filePath string)
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
	audioTrack, err := webrtc.NewTrackLocalStaticSample(webrtc.RTPCodecCapability{MimeType: webrtc.MimeTypeOpus}, fmt.Sprintf("audio_opus_%s", channelId), fmt.Sprintf("audio_opus_%s", channelId))
	if err != nil {
		return nil, err
	}
	rtpSender, err := peerConnection.AddTrack(audioTrack)
	if err != nil {
		return nil, err
	}

	// Read incoming RTCP packets
	// Before these packets are returned they are processed by interceptors. For things
	// like NACK this needs to be called.
	go func() {
		rtcpBuf := make([]byte, 1500)
		for {
			if _, _, rtcpErr := rtpSender.Read(rtcpBuf); rtcpErr != nil {
				return
			}
		}
	}()

	// save to store
	rtcConnection := &callRTCConn{
		peer:                    peerConnection,
		channelId:               channelId,
		receiverId:              receiverId,
		callerId:                c.botId,
		acceptCallAudioFile:     c.acceptCallAudioFile,
		exitCallAudioFile:       c.exitCallAudioFile,
		checkinSuccessAudioFile: c.checkinSuccessAudioFile,
		checkinFailAudioFile:    c.checkinFailAudioFile,
		audioTrack:              audioTrack,
		snapShootCount:          c.snapShootCount,
		rtpChan:                 make(chan *rtp.Packet),
		isVideoCall:             false,
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
					c.onICEConnectionStateChange(webrtc.ICEConnectionStateClosed, channelId, receiverId)
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

			unzipData, _ := utils.GzipUncompress(eventData.JsonData)
			var offer webrtc.SessionDescription
			err := json.Unmarshal([]byte(unzipData), &offer)
			if err != nil {
				return err
			}

			// only video call
			parsedSDP, err := offer.Unmarshal()
			if err != nil {
				return err
			}

			if !ok {
				var err error
				rtcConn, err = c.newCallRTCConnection(eventData.ChannelId, eventData.CallerId)
				if err != nil {
					return err
				}
			}

			for _, media := range parsedSDP.MediaDescriptions {
				if media.MediaName.Media == webrtc.RTPCodecTypeVideo.String() {
					rtcConn.isVideoCall = true
					break
				}
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

		case constants.WEBRTC_SDP_QUIT:
			c.onICEConnectionStateChange(webrtc.ICEConnectionStateClosed, eventData.ChannelId, eventData.ReceiverId)
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
		if rtcConn.(*callRTCConn).isVideoCall {
			rtcConn.(*callRTCConn).sendAudioTrack(c.acceptCallAudioFile)
			rtcConn.(*callRTCConn).saveTrackToImage(c.onImage, receiverId)
		} else {
			rtcConn.(*callRTCConn).sendAudioTrack(c.exitCallAudioFile)
			c.onICEConnectionStateChange(webrtc.ICEConnectionStateClosed, channelId, receiverId)
		}

	case webrtc.ICEConnectionStateClosed:
		if rtcConn.(*callRTCConn).peer.ConnectionState() != webrtc.PeerConnectionStateClosed {
			rtcConn.(*callRTCConn).peer.Close()
		}

		mapCallRtcConn.Delete(channelId)
	}
}

func (c *callService) SetAcceptCallFileAudio(filePath string) {
	c.acceptCallAudioFile = filePath
}

func (c *callService) SetExitCallFileAudio(filePath string) {
	c.exitCallAudioFile = filePath
}

func (c *callService) SetCheckinSuccessFileAudio(filePath string) {
	c.checkinSuccessAudioFile = filePath
}

func (c *callService) SetCheckinFailFileAudio(filePath string) {
	c.checkinFailAudioFile = filePath
}

func (c *callService) SetOnImage(onImage func(imgBase64 string) error, snapShootCount int) {
	c.snapShootCount = snapShootCount
	c.onImage = onImage
}

func (c *callRTCConn) sendAudioTrack(filePath string) error {

	// Open a OGG file and start reading using our OGGReader
	file, oggErr := os.Open(filePath)
	if oggErr != nil {
		return oggErr
	}

	// Open on oggfile in non-checksum mode.
	ogg, _, oggErr := oggreader.NewWith(file)
	if oggErr != nil {
		return oggErr
	}

	// Keep track of last granule, the difference is the amount of samples in the buffer
	var lastGranule uint64

	// It is important to use a time.Ticker instead of time.Sleep because
	// * avoids accumulating skew, just calling time.Sleep didn't compensate for the time spent parsing the data
	// * works around latency issues with Sleep (see https://github.com/golang/go/issues/44343)
	ticker := time.NewTicker(20 * time.Millisecond)
	defer ticker.Stop()
	for ; true; <-ticker.C {
		pageData, pageHeader, oggErr := ogg.ParseNextPage()
		if errors.Is(oggErr, io.EOF) {
			log.Println("All audio pages parsed and sent")
			return nil
		}

		if oggErr != nil {
			return oggErr
		}

		// The amount of samples is the difference between the last and current timestamp
		sampleCount := float64(pageHeader.GranulePosition - lastGranule)
		lastGranule = pageHeader.GranulePosition
		sampleDuration := time.Duration((sampleCount/48000)*1000) * time.Millisecond

		if oggErr = c.audioTrack.WriteSample(media.Sample{Data: pageData, Duration: sampleDuration}); oggErr != nil {
			return oggErr
		}
	}
	return nil

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
			c.sendAudioTrack(c.checkinSuccessAudioFile)

			close(c.rtpChan)
			c.peer.Close()
			return err
		}

		sampleBuilder = samplebuilder.New(20, &codecs.VP8Packet{}, 90000)
		decoder = vp8.NewDecoder()

		if imageCount > c.snapShootCount {
			c.sendAudioTrack(c.checkinFailAudioFile)

			close(c.rtpChan)
			c.peer.Close()
			return nil
		}
	}
}
