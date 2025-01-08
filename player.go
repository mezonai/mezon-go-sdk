package mezonsdk

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"

	"os"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/nccasia/mezon-go-sdk/utils"
	"github.com/pion/webrtc/v4"
	"github.com/pion/webrtc/v4/pkg/media"
	"github.com/pion/webrtc/v4/pkg/media/oggreader"
)

var (
	MapStreamingRtcConn sync.Map // map[channelId]*RTCConnection
)

type WsMsg struct {
	Key         string
	ClanId      string
	ChannelId   string
	UserId      string
	Username    string
	ClientId    string
	IsPublisher bool
	State       int
	Value       json.RawMessage
}

type Audience struct {
	userId string
	peer   *webrtc.PeerConnection
}

type streamingRTCConn struct {
	clanId    string
	channelId string
	username  string
	token     string

	// TODO: streaming video (#rapchieuphim)
	// videoTrack *webrtc.TrackLocalStaticRTP
	audioTrack *webrtc.TrackLocalStaticSample
	audiences  map[string]*Audience

	ctx        context.Context
	cancelFunc context.CancelFunc

	wsconn *websocket.Conn
}

var config = webrtc.Configuration{
	ICEServers: []webrtc.ICEServer{
		{
			URLs:       []string{"turn:turn.mezon.vn:5349", "stun:stun.l.google.com:19302"},
			Username:   "turnmezon",
			Credential: "QuTs4zUEcbylWemXL7MK",
		},
	},
}

type AudioPlayer interface {
	Play(filePath string) error
	Close(channelId string)
	Cancel(channelId string)
}

func NewAudioPlayer(clanId, channelId, userId, username, token string) (AudioPlayer, error) {
	// // Create a video track
	// videoTrack, err := webrtc.NewTrackLocalStaticRTP(webrtc.RTPCodecCapability{MimeType: webrtc.MimeTypeVP8}, fmt.Sprintf("video_vp8_%s", channelId), fmt.Sprintf("video_vp8_%s", channelId))
	// if err != nil {
	// 	return nil, err
	// }
	// _, err = peerConnection.AddTrack(videoTrack)
	// if err != nil {
	// 	return nil, err
	// }

	// Create a audio track
	audioTrack, err := webrtc.NewTrackLocalStaticSample(webrtc.RTPCodecCapability{MimeType: webrtc.MimeTypeOpus}, fmt.Sprintf("audio_opus_%s", channelId), fmt.Sprintf("audio_opus_%s", channelId))
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithCancel(context.Background())

	// save to store
	rtcConnection := &streamingRTCConn{
		clanId:     clanId,
		username:   username,
		token:      token,
		channelId:  channelId,
		audioTrack: audioTrack,
		ctx:        ctx,
		cancelFunc: cancel,
		audiences:  make(map[string]*Audience),
	}

	MapStreamingRtcConn.Store(channelId, rtcConnection)

	rtcConnection.audioTrack = audioTrack

	return rtcConnection, nil
}

func (s *streamingRTCConn) Close(channelId string) {
	rtcConn, ok := MapStreamingRtcConn.Load(channelId)
	if !ok {
		return
	}

	for _, audience := range rtcConn.(*streamingRTCConn).audiences {
		if audience != nil && audience.peer != nil && audience.peer.ConnectionState() != webrtc.PeerConnectionStateClosed {
			audience.peer.Close()
		}
	}
	rtcConn.(*streamingRTCConn).cancel()

	MapStreamingRtcConn.Delete(channelId)
}

func (s *streamingRTCConn) Cancel(channelId string) {
	s.sendMessage(&WsMsg{
		Key:       "stop_publisher",
		ClanId:    s.clanId,
		ChannelId: s.channelId,
	})
	s.cancel()
}

func (s *streamingRTCConn) cancel() {
	if s.cancelFunc != nil {
		s.cancelFunc()
	}

	s.ctx, s.cancelFunc = context.WithCancel(context.Background())
}

func (s *streamingRTCConn) OnWebsocketEvent(event *WsMsg) error {
	switch event.Key {
	case "session_subscriber":
		// receive offer from subscriber
		var offer webrtc.SessionDescription
		err := json.Unmarshal(event.Value, &offer)
		if err != nil {
			log.Println("unmarshal err: ", err)
			return err
		}
		_, err = s.createPeerConnection(&offer, event.UserId, event.ClientId)
		if err != nil {
			log.Println("session subscriber err: ", err)
			return err
		}
	case "ice_candidate":
		var i webrtc.ICECandidateInit
		err := json.Unmarshal(event.Value, &i)
		if err != nil {
			return err
		}

		return s.addICECandidate(i, event.ClientId)
	}

	return nil
}

func (s *streamingRTCConn) sendAnswer(clientId string) error {
	audience, ok := s.audiences[clientId]
	if !ok || audience.peer == nil {
		return errors.New("not init audience")
	}
	answer, err := audience.peer.CreateAnswer(nil)
	if err != nil {
		return err
	}
	if err := audience.peer.SetLocalDescription(answer); err != nil {
		return err
	}

	byteJson, _ := json.Marshal(answer.SDP)

	return s.sendMessage(&WsMsg{
		Key:       "sd_answer",
		ClanId:    s.clanId,
		ChannelId: s.channelId,
		ClientId:  clientId,
		Value:     byteJson,
	})
}

func (s *streamingRTCConn) onICECandidate(i *webrtc.ICECandidate, clanId, channelId, clientId string) error {
	if i == nil {
		return nil
	}
	// If you are serializing a candidate make sure to use ToJSON
	// Using Marshal will result in errors around `sdpMid`
	candidateString, err := json.Marshal(i.ToJSON())
	if err != nil {
		log.Println("onICECandidate err: ", err)
		return err
	}

	return s.sendMessage(&WsMsg{
		Key:       "ice_candidate",
		Value:     candidateString,
		ClanId:    clanId,
		ChannelId: channelId,
		ClientId:  clientId,
	})
}

func (s *streamingRTCConn) addICECandidate(i webrtc.ICECandidateInit, clientId string) error {
	var err error
	if audience, ok := s.audiences[clientId]; ok && audience.peer != nil {
		err = audience.peer.AddICECandidate(i)
	}
	return err
}

func (s *streamingRTCConn) Play(filePath string) error {
	basePath := "stn.mezon.vn"
	insecureSkip := true
	useSSL := true
	var dialer *websocket.Dialer
	if insecureSkip {
		tlsConfig := &tls.Config{
			InsecureSkipVerify: true,
		}
		dialer = &websocket.Dialer{
			TLSClientConfig: tlsConfig,
		}
	} else {
		dialer = websocket.DefaultDialer
	}
	basePath = utils.GetBasePath("ws", basePath, useSSL)

	conn, _, err := dialer.Dial(fmt.Sprintf("%s/ws?username=%s&token=%s", basePath, s.username, s.token), nil)
	if err != nil {
		log.Println("WebSocket connection open err: ", err)
		return err
	}

	s.wsconn = conn
	s.connectPublisher()

	go func() {
		for {
			msgType, databytes, err := conn.ReadMessage()
			if err != nil {
				log.Println("read error", err)
				if websocket.IsCloseError(err, websocket.CloseNormalClosure, websocket.CloseGoingAway) ||
					websocket.IsUnexpectedCloseError(err) {
					log.Println("WebSocket connection closed:", err)
					return
				}
				continue
			}

			if msgType != websocket.TextMessage {
				log.Println("unknown message type: ", msgType)
				continue
			}

			var msg WsMsg
			err = json.Unmarshal(databytes, &msg)
			if err != nil {
				log.Println("can't unmarshal json data: ", string(databytes))
				continue
			}

			if err := s.OnWebsocketEvent(&msg); err != nil {
				log.Println("on message error: ", err.Error())
				continue
			}
		}
	}()

	// Open a OGG file and start reading using our OGGReader
	file, oggErr := os.Open(filePath)
	if oggErr != nil {
		return oggErr
	}
	defer file.Close()

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

	for {
		select {
		case <-s.ctx.Done():
			return nil
		case <-ticker.C:
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

			if oggErr = s.audioTrack.WriteSample(media.Sample{Data: pageData, Duration: sampleDuration}); oggErr != nil {
				return oggErr
			}
		}
	}
}

func (s *streamingRTCConn) createPeerConnection(offer *webrtc.SessionDescription, userId, clientId string) (*webrtc.PeerConnection, error) {
	log.Printf("createPeerConnection %s \n", clientId)
	pc, err := webrtc.NewPeerConnection(config)
	if err != nil {
		return nil, err
	}
	if err := pc.SetRemoteDescription(*offer); err != nil {
		return nil, err
	}
	rtpSender, err := pc.AddTrack(s.audioTrack)
	if err != nil {
		return nil, err
	}

	s.audiences[clientId] = &Audience{
		peer:   pc,
		userId: userId,
	}
	s.sendAnswer(clientId)

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

	pc.OnICEConnectionStateChange(func(state webrtc.ICEConnectionState) {
		log.Printf("Connection State has changed %s \n", state.String())

		switch state {
		case webrtc.ICEConnectionStateDisconnected, webrtc.ICEConnectionStateClosed:
			_, ok := s.audiences[clientId]
			if ok {
				delete(s.audiences, clientId)
			}

			// goto sendMessage
			fallthrough

		case webrtc.ICEConnectionStateConnected:
			for clientId, audience := range s.audiences {
				if audience.peer == pc {
					s.sendMessage(&WsMsg{
						ClanId:    s.clanId,
						ChannelId: s.channelId,
						Key:       "session_state_changed",
						ClientId:  clientId,
						UserId:    audience.userId,
						State:     int(state),
					})
				}
			}
		}
	})

	pc.OnICECandidate(func(i *webrtc.ICECandidate) {
		s.onICECandidate(i, s.channelId, s.clanId, clientId)
	})

	return pc, nil
}

func (s *streamingRTCConn) connectPublisher() {
	err := s.sendMessage(&WsMsg{
		Key:       "connect_publisher",
		ClanId:    s.clanId,
		ChannelId: s.channelId,
		UserId:    s.username,
	})
	if err != nil {
		log.Println("can not send connect_publisher err: ", err)
	}
}

func (s *streamingRTCConn) sendMessage(data *WsMsg) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Println("can not marshal json err: ", err)
		return err
	}

	return s.wsconn.WriteMessage(websocket.TextMessage, jsonData)
}
