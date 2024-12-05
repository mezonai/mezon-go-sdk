package client

import (
	"crypto/tls"
	"log"
	"mezon-sdk/constants"
	"mezon-sdk/mezon-protobuf/mezon/v2/common/rtapi"
	"time"

	"github.com/gorilla/websocket"
	"google.golang.org/protobuf/proto"
)

var (
	client *WSConnection
)

type WSConnection struct {
	conn        *websocket.Conn
	baseWS      string
	clanId      string
	recvHandler []func(*rtapi.Envelope) error
}

type IWSConnection interface {
	SendMessage(data *rtapi.Envelope) error
	SetRecvHandler(recvHandler func(*rtapi.Envelope) error)
	Close() error
}

func GetWSConnection(baseWS string, clanId string, recvHandler ...func(*rtapi.Envelope) error) (IWSConnection, error) {
	if client == nil {
		client = &WSConnection{
			baseWS:      baseWS,
			clanId:      baseWS,
			recvHandler: recvHandler,
		}
		if err := client.newWSConnection(); err != nil {
			return nil, err
		}
	}

	return client, nil
}

func (s *WSConnection) newWSConnection() error {
	if s.baseWS == "" {
		s.baseWS = constants.DefaultBaseWS
	}

	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
	}
	dialer := websocket.Dialer{
		TLSClientConfig: tlsConfig,
	}
	conn, _, err := dialer.Dial(s.baseWS, nil)
	if err != nil {
		return err
	}

	client.conn = conn
	if err = client.joinClan(s.clanId); err != nil {
		return err
	}

	client.reconnect()
	client.pingPong()
	client.recvMessage()

	return nil
}

func (s *WSConnection) Close() error {
	return s.conn.Close()
}

func (s *WSConnection) SendMessage(data *rtapi.Envelope) error {
	jsonData, err := proto.Marshal(data)
	if err != nil {
		return err
	}
	return s.conn.WriteMessage(websocket.BinaryMessage, jsonData)
}

func (s *WSConnection) SetRecvHandler(recvHandler func(*rtapi.Envelope) error) {
	s.recvHandler = append(s.recvHandler, recvHandler)
}

func (s *WSConnection) joinClan(clanId string) error {
	clanJoinData := &rtapi.Envelope{
		Cid: "",
		Message: &rtapi.Envelope_ClanJoin{
			ClanJoin: &rtapi.ClanJoin{
				ClanId: clanId,
			},
		},
	}
	return s.SendMessage(clanJoinData)
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

			for _, handler := range s.recvHandler {
				err = handler(request)
				if err != nil {
					continue
				}
			}
		}
	}()
}
