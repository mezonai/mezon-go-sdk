package main

import (
	mezonsdk "mezon-go-sdk"

	"github.com/pion/webrtc/v4"
)

func main() {

	// streaming
	Streaming()

	select {}
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
