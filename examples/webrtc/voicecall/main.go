package main

import (
	"fmt"
	"time"

	mezonsdk "github.com/nccasia/mezon-go-sdk"

	"github.com/pion/webrtc/v4"
)

func main() {

	// call
	Call()

	select {}
}

func Call() {
	channelId := "1840660756006178816"
	conn, err := mezonsdk.NewWSConnection(&mezonsdk.Config{
		BasePath: "dev-mezon.nccsoft.vn:7305",
		// BasePath:     "api.mezon.vn",
		ApiKey:       "7663586b614xxxxxx75356a5a4d52",
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
				Credential:     "xxxxx",
				CredentialType: webrtc.ICECredentialTypePassword,
			},
		},
	}, "images", 10)

	conn.SetOnWebrtcSignalingFwd(callService.OnWebsocketEvent)

	loop := true
	for loop {
		if callService.GetRTCConnectionState(channelId) == webrtc.PeerConnectionStateConnected {
			err = callService.SendFile(channelId, "file:///home/minhnv/Music/test.mp3")
			if err != nil {
				panic(err)
			}

			fmt.Println("webrtc send file")
			break
		}

		if callService.GetRTCConnectionState(channelId) == webrtc.PeerConnectionStateFailed {
			loop = false
		}

		time.Sleep(500 * time.Millisecond)
	}
}
