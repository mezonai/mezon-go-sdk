package mezonsdk

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"image/jpeg"
	"log"
	"mezon-sdk/constants"
	"mezon-sdk/mezon-protobuf/mezon/v2/common/rtapi"
	"mezon-sdk/utils"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/go-gst/go-gst/gst"
	"github.com/go-gst/go-gst/gst/app"
	"github.com/pion/rtcp"
	"github.com/pion/rtp"
	"github.com/pion/rtp/codecs"
	"github.com/pion/webrtc/v4"
	"github.com/pion/webrtc/v4/pkg/media/samplebuilder"
	"golang.org/x/image/vp8"
)

func init() {
	// Initialize GStreamer
	gst.Init(nil)
}

var mapCallRtcConn sync.Map // map[channelId]*RTCConnection

type callRTCConn struct {
	peer       *webrtc.PeerConnection
	channelId  string
	receiverId string
	callerId   string

	audioTrack    *webrtc.TrackLocalStaticRTP
	rtpChan       chan *rtp.Packet
	pathImage     string
	timeSaveImage int
}

type callService struct {
	botId         string
	ws            IWSConnection
	config        webrtc.Configuration
	pathImage     string
	timeSaveImage int
}

type ICallService interface {
	SendFile(channelId, filePath string) error
	OnWebsocketEvent(event *rtapi.Envelope) error
	GetRTCConnectionState(channelId string) webrtc.PeerConnectionState
}

func NewCallService(botId string, wsConn IWSConnection, config webrtc.Configuration, pathImage string, timeSaveImage int) ICallService {
	return &callService{
		botId:         botId,
		ws:            wsConn,
		config:        config,
		pathImage:     pathImage,
		timeSaveImage: timeSaveImage,
	}
}

func (c *callService) newCallRTCConnection(channelId, receiverId string) (*callRTCConn, error) {

	peerConnection, err := webrtc.NewPeerConnection(c.config)
	if err != nil {
		return nil, err
	}

	// Create a audio track
	audioTrack, err := webrtc.NewTrackLocalStaticRTP(webrtc.RTPCodecCapability{MimeType: webrtc.MimeTypeOpus}, fmt.Sprintf("audio_opus_%s", channelId), fmt.Sprintf("audio_opus_%s", channelId))
	if err != nil {
		return nil, err
	}
	_, err = peerConnection.AddTrack(audioTrack)
	if err != nil {
		return nil, err
	}

	// save to store
	rtcConnection := &callRTCConn{
		peer:          peerConnection,
		channelId:     channelId,
		receiverId:    receiverId,
		callerId:      c.botId,
		audioTrack:    audioTrack,
		pathImage:     c.pathImage,
		timeSaveImage: c.timeSaveImage,
		rtpChan:       make(chan *rtp.Packet),
	}
	mapCallRtcConn.Store(channelId, rtcConnection)

	peerConnection.OnICEConnectionStateChange(func(state webrtc.ICEConnectionState) {
		c.onICEConnectionStateChange(state, channelId)
	})
	peerConnection.OnICECandidate(func(i *webrtc.ICECandidate) {
		c.onICECandidate(i, channelId, c.botId, receiverId)
	})

	peerConnection.OnTrack(func(track *webrtc.TrackRemote, receiver *webrtc.RTPReceiver) {

		// save image by time receive track
		if rtcConnection.timeSaveImage > 0 && track.Kind() == webrtc.RTPCodecTypeVideo {
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

			startTime := time.Now()
			for {
				// Read RTP Packets in a loop
				rtpPacket, _, readErr := track.ReadRTP()
				if readErr != nil {
					log.Printf("track read rtp error: %+v \n", readErr)
					return
				}

				if time.Since(startTime) > time.Duration(c.timeSaveImage)*time.Second {
					break
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

	go rtcConnection.saveTrackToImage(receiverId)

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

func (c *callService) onICEConnectionStateChange(state webrtc.ICEConnectionState, channelId string) {
	log.Printf("Connection State has changed %s \n", state.String())

	switch state {
	case webrtc.ICEConnectionStateConnected:

	case webrtc.ICEConnectionStateClosed:
		rtcConn, ok := mapCallRtcConn.Load(channelId)
		if !ok {
			return
		}

		if rtcConn.(*callRTCConn).peer == nil {
			return
		}

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

// Create the appropriate GStreamer pipeline depending on what codec we are working with
func (c *callRTCConn) pipelineForCodec(codecName string, tracks []*webrtc.TrackLocalStaticRTP, pipelineSrc string) {
	log.Printf("[pipelineForCodec] codecName: %s, len(track): %d, pipelineSrc: %s \n", codecName, len(tracks), pipelineSrc)

	pipelineStr := "appsink name=appsink"
	switch codecName {
	case "vp8":
		pipelineStr = pipelineSrc + " ! vp8enc error-resilient=partitions keyframe-max-dist=10 auto-alt-ref=true cpu-used=5 deadline=1 ! " + pipelineStr
	case "vp9":
		pipelineStr = pipelineSrc + " ! vp9enc ! " + pipelineStr
	case "h264":
		pipelineStr = pipelineSrc + " ! video/x-raw,format=I420 ! x264enc speed-preset=ultrafast tune=zerolatency key-int-max=20 ! video/x-h264,stream-format=byte-stream ! " + pipelineStr
	case "opus":
		pipelineStr = pipelineSrc + " ! opusenc ! " + pipelineStr
	case "pcmu":
		pipelineStr = pipelineSrc + " ! audio/x-raw, rate=8000 ! mulawenc ! " + pipelineStr
	case "pcma":
		pipelineStr = pipelineSrc + " ! audio/x-raw, rate=8000 ! alawenc ! " + pipelineStr
	case "mp3":
		// Opus pipeline with RTP payloading
		pipelineStr = pipelineSrc + " ! audioresample ! audio/x-raw,rate=48000 ! opusenc ! rtpopuspay ! " + pipelineStr
	default:
		log.Println("Unhandled codec " + codecName)
		return
	}

	pipeline, err := gst.NewPipelineFromString(pipelineStr)
	if err != nil {
		log.Println("new pipeline gstreamer error: ", err)
		return
	}

	if err = pipeline.SetState(gst.StatePlaying); err != nil {
		log.Println("pipeline gstreamer set state error: ", err)
		return
	}

	appSink, err := pipeline.GetElementByName("appsink")
	if err != nil {
		log.Println("pipeline gstreamer set state error: ", err)
		return
	}

	app.SinkFromElement(appSink).SetCallbacks(&app.SinkCallbacks{
		NewSampleFunc: func(sink *app.Sink) gst.FlowReturn {
			sample := sink.PullSample()
			if sample == nil {
				return gst.FlowEOS
			}

			buffer := sample.GetBuffer()
			if buffer == nil {
				return gst.FlowError
			}

			samples := buffer.Map(gst.MapRead).Bytes()
			defer buffer.Unmap()

			for _, t := range tracks {
				if _, err := t.Write(samples); err != nil {
					log.Println("track write error: ", err)
					return gst.FlowError
				}
			}

			return gst.FlowOK
		},
	})
}

func (c *callRTCConn) saveTrackToImage(receiverId string) error {
	log.Printf("[saveTrackToImage] receiverId: %s \n", receiverId)

	currentDate := time.Now().Format("2006-01-02")
	currentTime := time.Now().Format("150405")
	folderPath := filepath.Join(".", c.pathImage, currentDate)
	if err := os.MkdirAll(folderPath, os.ModePerm); err != nil {
		return err
	}

	imagePath := filepath.Join(folderPath, fmt.Sprintf("%s_%s.jpg", currentTime, receiverId))
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

		file, err := os.Create(imagePath)
		if err != nil {
			log.Printf("failed to create image file: %+v \n", err)
			return fmt.Errorf("failed to create image file: %w", err)
		}

		// Encode to (RGB) jpeg
		if err = jpeg.Encode(file, img, nil); err != nil {
			log.Printf("jpeg.Encode error: %+v \n", err)
			return err
		}
		file.Close()
	}
}
