package main

import (
	"fmt"
	mezonsdk "mezon-sdk"
	"time"

	"github.com/pion/webrtc/v4"
)

func main() {

	// streaming
	// Streaming()

	// call
	Call()

	select {}
}

func Call() {
	channelId := "1840660756006178816"
	conn, err := mezonsdk.NewWSConnection(&mezonsdk.Config{
		BasePath: "dev-mezon.nccsoft.vn:7305",
		// BasePath:     "api.mezon.vn",
		ApiKey:       "7663586b614xxxxx47175356a5a4d52",
		Timeout:      10,
		InsecureSkip: true,
		UseSSL:       true,
	}, "")
	if err != nil {
		panic(err)
	}
	fmt.Println("ws connected")

	callService := mezonsdk.NewCallService("1835944289075466240", conn, webrtc.Configuration{
		ICEServers: []webrtc.ICEServer{
			{
				URLs:           []string{"turn:turn.mezon.vn:5349"},
				Username:       "turnmezon",
				Credential:     "QuTs4zUEcbylWemXL7MK",
				CredentialType: webrtc.ICECredentialTypePassword,
			},
		},
	}, "images")

	conn.SetOnWebrtcSignalingFwd(callService.OnWebsocketEvent)

	for true {
		if callService.GetRTCConnectionState(channelId) == webrtc.PeerConnectionStateConnected {
			err = callService.SendFile(channelId, "file:///home/minhnv/Music/test.mp3")
			if err != nil {
				panic(err)
			}

			fmt.Println("webrtc send file")
			break
		}

		time.Sleep(1 * time.Second)
	}
}

func Streaming() {

	conn, err := mezonsdk.NewWSConnection(&mezonsdk.Config{
		BasePath: "dev-mezon.nccsoft.vn:7305",
		// BasePath:     "api.mezon.vn",
		ApiKey:       "7663586b61xxxxxxxx356a5a4d52",
		Timeout:      10,
		InsecureSkip: true,
		UseSSL:       true,
	}, "1827955317304987648")
	if err != nil {
		panic(err)
	}

	filePath := ""
	rtcConn, err := mezonsdk.NewStreamingRTCConnection(webrtc.Configuration{
		ICEServers: []webrtc.ICEServer{
			{
				URLs: []string{"stun.l.google.com:19302"},
			},
		},
	}, conn, "1827987498463137792", "1827987498463137792")
	if err != nil {
		panic(err)
	}

	err = rtcConn.SendFile(filePath)
	if err != nil {
		panic(err)
	}
}
