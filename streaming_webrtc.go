package mezonsdk

import (
	"encoding/json"
	"fmt"
	"log"
	"sync"

	"github.com/nccasia/mezon-go-sdk/mezon-protobuf/mezon/v2/common/rtapi"

	"github.com/nccasia/mezon-go-sdk/utils"

	"github.com/nccasia/mezon-go-sdk/constants"

	"github.com/pion/webrtc/v4"
)

var (
	mapStreamingRtcConn sync.Map // map[channelId]*RTCConnection
)

type streamingRTCConn struct {
	peer *webrtc.PeerConnection
	ws   IWSConnection

	clanId    string
	channelId string

	videoTrack *webrtc.TrackLocalStaticRTP
	audioTrack *webrtc.TrackLocalStaticRTP
}

type IStreamingRTCConnection interface {
	SendFile(filePath string) error
	Close(channelId string)
}

func NewStreamingRTCConnection(config webrtc.Configuration, wsConn IWSConnection, clanId, channelId string) (IStreamingRTCConnection, error) {
	peerConnection, err := webrtc.NewPeerConnection(config)
	if err != nil {
		return nil, err
	}

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
	rtcConnection := &streamingRTCConn{
		peer:       peerConnection,
		ws:         wsConn,
		clanId:     clanId,
		channelId:  channelId,
		videoTrack: videoTrack,
		audioTrack: audioTrack,
	}
	wsConn.SetOnJoinStreamingChannel(rtcConnection.onWebsocketEvent)
	mapStreamingRtcConn.Store(channelId, rtcConnection)

	peerConnection.OnICEConnectionStateChange(func(state webrtc.ICEConnectionState) {
		log.Printf("Connection State has changed %s \n", state.String())

		switch state {
		case webrtc.ICEConnectionStateConnected:
			// TODO: event ice connected
		case webrtc.ICEConnectionStateClosed:
			rtcConn, ok := mapStreamingRtcConn.Load(channelId)
			if !ok {
				return
			}

			if rtcConn.(*streamingRTCConn).peer == nil {
				return
			}

			if rtcConn.(*streamingRTCConn).peer.ConnectionState() != webrtc.PeerConnectionStateClosed {
				rtcConn.(*streamingRTCConn).peer.Close()
			}

			mapStreamingRtcConn.Delete(channelId)
		}
	})
	peerConnection.OnICECandidate(func(i *webrtc.ICECandidate) {
		rtcConnection.onICECandidate(i, channelId, clanId)
	})

	// send offer
	rtcConnection.sendOffer()

	return rtcConnection, nil
}

func (c *streamingRTCConn) SendFile(filePath string) error {
	// Start pushing buffers on these tracks
	audioSrc := fmt.Sprintf("uridecodebin uri=%s ! queue ! audioconvert", filePath)
	videoSrc := fmt.Sprintf("uridecodebin uri=%s ! videoscale ! video/x-raw, width=320, height=240 ! queue ", filePath)

	c.pipelineForCodec("opus", []*webrtc.TrackLocalStaticRTP{c.audioTrack}, audioSrc)
	c.pipelineForCodec("vp8", []*webrtc.TrackLocalStaticRTP{c.videoTrack}, videoSrc)

	return nil
}

func (c *streamingRTCConn) Close(channelId string) {
	rtcConn, ok := mapStreamingRtcConn.Load(channelId)
	if !ok {
		return
	}

	if rtcConn.(*streamingRTCConn).peer == nil {
		return
	}

	if rtcConn.(*streamingRTCConn).peer.ConnectionState() != webrtc.PeerConnectionStateClosed {
		rtcConn.(*streamingRTCConn).peer.Close()
	}

	mapStreamingRtcConn.Delete(channelId)
}

func (c *streamingRTCConn) onWebsocketEvent(event *rtapi.Envelope) error {
	switch event.Message.(type) {
	case *rtapi.Envelope_JoinStreamingChannel:
		eventData := event.GetJoinStreamingChannel()
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

func (c *streamingRTCConn) sendOffer() error {
	offer, err := c.peer.CreateOffer(nil)
	if err != nil {
		return err
	}
	if err := c.peer.SetLocalDescription(offer); err != nil {
		return err
	}

	byteJson, _ := json.Marshal(offer)
	dataEnc, _ := utils.GzipCompress(string(byteJson))

	// send socket signaling, gzip compress data
	return c.ws.SendMessage(&rtapi.Envelope{
		Cid: "",
		Message: &rtapi.Envelope_JoinStreamingChannel{
			JoinStreamingChannel: &rtapi.JoinStreamingChannel{
				JsonData:      dataEnc,
				DataType:      constants.WEBRTC_SDP_OFFER,
				StreamingPush: true,
				ClanId:        c.clanId,
				ChannelId:     c.channelId,
			},
		},
	})
}

func (c *streamingRTCConn) onICECandidate(i *webrtc.ICECandidate, clanId, channelId string) error {
	if i == nil {
		return nil
	}
	// If you are serializing a candidate make sure to use ToJSON
	// Using Marshal will result in errors around `sdpMid`
	candidateString, err := json.Marshal(i.ToJSON())
	if err != nil {
		return err
	}

	return c.ws.SendMessage(&rtapi.Envelope{Message: &rtapi.Envelope_JoinStreamingChannel{JoinStreamingChannel: &rtapi.JoinStreamingChannel{
		DataType:      constants.WEBRTC_ICE_CANDIDATE,
		JsonData:      string(candidateString),
		ChannelId:     channelId,
		ClanId:        clanId,
		StreamingPush: true,
	}}})
}

func (c *streamingRTCConn) addICECandidate(i webrtc.ICECandidateInit) error {
	return c.peer.AddICECandidate(i)
}

func (c *streamingRTCConn) pipelineForCodec(codecName string, tracks []*webrtc.TrackLocalStaticRTP, pipelineSrc string) {
	//TODO: ffmpeg or convert file in bot repo
}
