package mezonsdk

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"sync"

	"github.com/gorilla/websocket"
	"github.com/nccasia/mezon-go-sdk/constants"
	"github.com/nccasia/mezon-go-sdk/utils"
	"github.com/pion/webrtc/v4"
)

var (
	MapStreamingRtcConn sync.Map // map[channelId]*RTCConnection
)

type FileUrl struct {
	FileUrl string
}

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
}

type streamingMediaConn struct {
	clanId    string
	channelId string
	userId    string
	username  string
	token     string

	audiences map[string]*Audience

	ctx        context.Context
	cancelFunc context.CancelFunc

	wsconn *websocket.Conn
}

var config = webrtc.Configuration{
	ICEServers: []webrtc.ICEServer{
		{
			URLs:       []string{"turn:replay.mezon.vn:5349", "stun:stun.l.google.com:19302"},
			Username:   "turnmezon",
			Credential: "QuTs4zUEcbylWemXL7MK",
		},
	},
}

type AudioPlayer interface {
	Play(fileLink string) error
	Close(channelId string)
	Cancel(channelId string)
}

func NewAudioPlayer(clanId, channelId, userId, username, token string) (AudioPlayer, error) {
	ctx, cancel := context.WithCancel(context.Background())

	// save to store
	rtcConnection := &streamingMediaConn{
		clanId:     clanId,
		channelId:  channelId,
		userId:     userId,
		username:   username,
		token:      token,
		ctx:        ctx,
		cancelFunc: cancel,
		audiences:  make(map[string]*Audience),
	}

	MapStreamingRtcConn.Store(channelId, rtcConnection)

	return rtcConnection, nil
}

func (s *streamingMediaConn) Close(channelId string) {
	rtcConn, ok := MapStreamingRtcConn.Load(channelId)
	if !ok {
		return
	}

	rtcConn.(*streamingMediaConn).cancel()
	s.stopPublisher()

	MapStreamingRtcConn.Delete(channelId)
}

func (s *streamingMediaConn) Cancel(channelId string) {
	s.cancel()
	s.stopPublisher()
}

func (s *streamingMediaConn) cancel() {
	if s.cancelFunc != nil {
		s.cancelFunc()
	}

	s.ctx, s.cancelFunc = context.WithCancel(context.Background())
}

func (s *streamingMediaConn) Play(fileUrl string) error {
	var dialer *websocket.Dialer
	if constants.InsecureSkip {
		tlsConfig := &tls.Config{
			InsecureSkipVerify: true,
		}
		dialer = &websocket.Dialer{
			TLSClientConfig: tlsConfig,
		}
	} else {
		dialer = websocket.DefaultDialer
	}

	basePath := utils.GetBasePath("ws", constants.StnBasePath, constants.UseSSL)

	conn, _, err := dialer.Dial(fmt.Sprintf("%s/ws?username=%s&token=%s", basePath, s.username, s.token), nil)
	if err != nil {
		log.Println("WebSocket connection open err: ", err)
		return err
	}

	s.wsconn = conn
	s.sessionPublisher()
	time.Sleep(100 * time.Millisecond)
	s.connectPublisher(fileUrl)

	return nil
}

func (s *streamingMediaConn) sendToStn(audience *Audience, state int, clientId string) {
	s.sendMessage(&WsMsg{
		ClanId:    s.clanId,
		ChannelId: s.channelId,
		Key:       "session_state_changed",
		ClientId:  clientId,
		UserId:    audience.userId,
		State:     state,
	})
}

func (s *streamingMediaConn) sessionPublisher() {
	err := s.sendMessage(&WsMsg{
		Key:       "session_publisher",
		ClanId:    s.clanId,
		ChannelId: s.channelId,
		UserId:    s.userId,
	})
	if err != nil {
		log.Println("can not send connect_publisher err: ", err)
	}
}

func (s *streamingMediaConn) connectPublisher(fileUrl string) {
	value, _ := json.Marshal(FileUrl{FileUrl: fileUrl})
	err := s.sendMessage(&WsMsg{
		Key:       "connect_publisher",
		ClanId:    s.clanId,
		ChannelId: s.channelId,
		UserId:    s.userId,
		Value:     value,
	})
	if err != nil {
		log.Println("can not send connect_publisher err: ", err)
	}
}

func (s *streamingMediaConn) stopPublisher() {
	err := s.sendMessage(&WsMsg{
		Key:       "stop_publisher",
		ClanId:    s.clanId,
		ChannelId: s.channelId,
		UserId:    s.userId,
	})
	if err != nil {
		log.Println("can not send stop_publisher err: ", err)
	}
}

func (s *streamingMediaConn) sendMessage(data *WsMsg) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Println("can not marshal json err: ", err)
		return err
	}

	return s.wsconn.WriteMessage(websocket.TextMessage, jsonData)
}
