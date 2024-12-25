package mezonsdk

import (
	"crypto/tls"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/nccasia/mezon-go-sdk/configs"
	"github.com/nccasia/mezon-go-sdk/utils"

	"github.com/nccasia/mezon-go-sdk/mezon-protobuf/mezon/v2/common/rtapi"

	"github.com/gorilla/websocket"
	"google.golang.org/protobuf/proto"
)

func recvDefaultHandler(e *rtapi.Envelope) error {
	return nil
}

type WSConnection struct {
	conn                   *websocket.Conn
	dialer                 *websocket.Dialer
	basePath               string
	token                  string
	clanIds                []string
	mu                     sync.Mutex
	onJoinStreamingChannel func(*rtapi.Envelope) error
	onWebrtcSignalingFwd   func(*rtapi.Envelope) error
	onPong                 func(*rtapi.Envelope) error
	onChannelMessage       func(*rtapi.Envelope) error
}

type IWSConnection interface {
	SendMessage(data *rtapi.Envelope) error
	SetOnJoinStreamingChannel(recvHandler func(*rtapi.Envelope) error)
	SetOnWebrtcSignalingFwd(recvHandler func(*rtapi.Envelope) error)
	SetOnPong(recvHandler func(*rtapi.Envelope) error)
	SetOnChannelMessage(recvHandler func(*rtapi.Envelope) error)
	Close() error
}

func NewWSConnection(c *configs.Config, token string) (IWSConnection, error) {
	socket := &WSConnection{
		token:                  token,
		basePath:               utils.GetBasePath("ws", c.BasePath, c.UseSSL),
		onJoinStreamingChannel: recvDefaultHandler,
		onWebrtcSignalingFwd:   recvDefaultHandler,
		onPong:                 recvDefaultHandler,
		onChannelMessage:       recvDefaultHandler,
	}

	if c.InsecureSkip {
		tlsConfig := &tls.Config{
			InsecureSkipVerify: true,
		}
		socket.dialer = &websocket.Dialer{
			TLSClientConfig: tlsConfig,
		}
	} else {
		socket.dialer = websocket.DefaultDialer
	}

	if err := socket.newWSConnection(); err != nil {
		return nil, err
	}

	// join DM for default
	socket.joinClan([]string{""})

	return socket, nil
}

func (s *WSConnection) newWSConnection() error {
	conn, _, err := s.dialer.Dial(fmt.Sprintf("%s/ws?lang=en&status=true&token=%s&format=protobuf", s.basePath, s.token), nil)
	if err != nil {
		log.Println("WebSocket connection open err: ", err)
		return err
	}

	s.conn = conn

	if err = s.joinClan(s.clanIds); err != nil {
		log.Printf("WebSocket join clan: %s, err: %+v \n", err)
		return err
	}

	s.reconnect()
	s.pingPong()
	s.recvMessage()

	return nil
}

func (s *WSConnection) Close() error {
	return s.conn.Close()
}

func (s *WSConnection) SendMessage(data *rtapi.Envelope) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	jsonData, err := proto.Marshal(data)
	if err != nil {
		return err
	}

	return s.conn.WriteMessage(websocket.BinaryMessage, jsonData)
}

func (s *WSConnection) joinClan(clanIds []string) error {
	for _, clanId := range clanIds {
		clanJoinData := &rtapi.Envelope{
			Cid: "",
			Message: &rtapi.Envelope_ClanJoin{
				ClanJoin: &rtapi.ClanJoin{
					ClanId: clanId,
				},
			},
		}
		s.SendMessage(clanJoinData)
	}

	return nil
}

func (s *WSConnection) pingPong() {
	// Ping Handler
	s.conn.SetPingHandler(func(appData string) error {
		if err := s.conn.WriteMessage(websocket.PongMessage, []byte(appData)); err != nil {
			return err
		}
		return nil
	})

	go func() {
		ticker := time.NewTicker(10 * time.Second)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				err := s.SendMessage(&rtapi.Envelope{
					Message: &rtapi.Envelope_Ping{
						Ping: &rtapi.Ping{},
					},
				})
				if err != nil {
					if websocket.IsCloseError(err, websocket.CloseNormalClosure, websocket.CloseGoingAway) ||
						websocket.IsUnexpectedCloseError(err) {
						log.Println("WebSocket connection closed:", err)
						return
					}
					continue
				}
			}
		}
	}()
}

func (s *WSConnection) reconnect() {
	s.conn.SetCloseHandler(func(code int, text string) error {
		log.Printf("WebSocket closed (code: %d, text: %s). Attempting to reconnect...", code, text)
		if s.conn != nil {
			_ = s.conn.Close()
		}
		return s.newWSConnection()
	})
}

func (s *WSConnection) recvMessage() {
	go func() {
		for {
			_, databytes, err := s.conn.ReadMessage()
			if err != nil {
				if websocket.IsCloseError(err, websocket.CloseNormalClosure, websocket.CloseGoingAway) ||
					websocket.IsUnexpectedCloseError(err) {
					log.Println("WebSocket connection closed:", err)
					return
				}
				continue
			}

			request := &rtapi.Envelope{}
			err = proto.Unmarshal(databytes, request)
			if err != nil {
				continue
			}

			if request.Cid != "" {
				continue
			}

			switch request.Message.(type) {
			case *rtapi.Envelope_JoinStreamingChannel:
				s.onJoinStreamingChannel(request)
			case *rtapi.Envelope_WebrtcSignalingFwd:
				s.onWebrtcSignalingFwd(request)
			case *rtapi.Envelope_Pong:
				s.onPong(request)
			case *rtapi.Envelope_ChannelMessage:
				s.onChannelMessage(request)
			}
		}
	}()
}

func (s *WSConnection) SetOnJoinStreamingChannel(recvHandler func(*rtapi.Envelope) error) {
	s.onJoinStreamingChannel = recvHandler
}

func (s *WSConnection) SetOnWebrtcSignalingFwd(recvHandler func(*rtapi.Envelope) error) {
	s.onWebrtcSignalingFwd = recvHandler
}

func (s *WSConnection) SetOnPong(recvHandler func(*rtapi.Envelope) error) {
	s.onPong = recvHandler
}

func (s *WSConnection) SetOnChannelMessage(recvHandler func(*rtapi.Envelope) error) {
	s.onChannelMessage = recvHandler
}
