package main

import (
	mezonsdk "github.com/nccasia/mezon-go-sdk"
	radiostation "github.com/nccasia/mezon-go-sdk/radio-station"
	"github.com/pion/webrtc/v4"
)

func main() {

	// streaming
	Streaming()

	select {}
}

func Streaming() {

	clanId := "xxx"
	channelId := "xxx"
	wsConn, err := radiostation.NewWSConnection(&radiostation.Config{
		BasePath:     "stn.nccsoft.vn",
		Timeout:      10,
		InsecureSkip: true,
		UseSSL:       true,
	}, clanId)
	if err != nil {
		panic(err)
	}

	filePath := "test.mp3"
	rtcConn, err := mezonsdk.NewStreamingRTCConnection(webrtc.Configuration{
		ICEServers: []webrtc.ICEServer{
			{
				URLs: []string{"stun.l.google.com:19302"}, // TODO: check radio station ice server
			},
		},
	}, wsConn, clanId, channelId)
	if err != nil {
		panic(err)
	}

	err = rtcConn.SendFile(filePath)
	if err != nil {
		panic(err)
	}
}
