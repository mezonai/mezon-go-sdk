package mezonsdk

import (
	"encoding/json"
	"fmt"
	"log"
	"mezon-sdk/mezon-protobuf/mezon/v2/common/rtapi"
	"sync"

	"mezon-sdk/constants"
	"mezon-sdk/utils"

	"github.com/go-gst/go-gst/gst"
	"github.com/go-gst/go-gst/gst/app"
	"github.com/pion/webrtc/v4"
)

var (
	mapChannelRtcconnection sync.Map // map[channelId]*RTCConnection
)

type RTCConnection struct {
	peer *webrtc.PeerConnection
	ws   IWSConnection

	videoTrack *webrtc.TrackLocalStaticRTP
	audioTrack *webrtc.TrackLocalStaticRTP
}

type IRTCConnection interface {
	SendFile(filePath string) error
}

func NewRTCConnection(config webrtc.Configuration, wsConn IWSConnection, clanId, channelId string) (IRTCConnection, error) {
	peerConnection, err := webrtc.NewPeerConnection(config)
	if err != nil {
		return nil, err
	}
	peerConnection.OnICEConnectionStateChange(func(connectionState webrtc.ICEConnectionState) {
		fmt.Printf("Connection State has changed %s \n", connectionState.String())
	})

	// Create a video track
	videoTrack, err := webrtc.NewTrackLocalStaticRTP(webrtc.RTPCodecCapability{MimeType: webrtc.MimeTypeVP8}, fmt.Sprintf("video_vp8_%s", channelId), fmt.Sprintf("video_vp8_%s", channelId))
	if err != nil {
		return nil, err
	}
	_, err = peerConnection.AddTrack(videoTrack)
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
	rtcConnection := &RTCConnection{
		peer: peerConnection,
		ws:   wsConn,
	}
	wsConn.SetRecvHandler(rtcConnection.onWebsocketEvent)
	mapChannelRtcconnection.Store(channelId, rtcConnection)

	// send offer
	rtcConnection.sendOffer()

	return rtcConnection, nil
}

func (c *RTCConnection) SendFile(filePath string) error {
	// Start pushing buffers on these tracks
	audioSrc := fmt.Sprintf("uridecodebin uri=%s ! queue ! audioconvert", filePath)
	videoSrc := fmt.Sprintf("uridecodebin uri=%s ! videoscale ! video/x-raw, width=320, height=240 ! queue ", filePath)

	c.pipelineForCodec("opus", []*webrtc.TrackLocalStaticRTP{c.audioTrack}, audioSrc)
	c.pipelineForCodec("vp8", []*webrtc.TrackLocalStaticRTP{c.videoTrack}, videoSrc)

	return nil
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

			return c.peer.SetRemoteDescription(answer)

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

// Create the appropriate GStreamer pipeline depending on what codec we are working with
func (c *RTCConnection) pipelineForCodec(codecName string, tracks []*webrtc.TrackLocalStaticRTP, pipelineSrc string) {
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
