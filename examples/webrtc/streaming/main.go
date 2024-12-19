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

	clanId := "3456110592"    // KOMU_2
	channelId := "4311748608" // NCC8
	userId := "4198400"       // longma350
	wsConn, err := radiostation.NewWSConnection(&radiostation.Config{
		BasePath:     "stn.nccsoft.vn",
		Timeout:      10,
		InsecureSkip: true,
		UseSSL:       true,
	}, clanId, channelId, userId)
	if err != nil {
		panic(err)
	}

	// ffmpeg -i test.mp3 -c:a libopus -page_duration 20000 test.ogg;
	filePath := "ncc8.ogg"
	rtcConn, err := mezonsdk.NewStreamingRTCConnection(webrtc.Configuration{
		ICEServers: []webrtc.ICEServer{
			{
				URLs: []string{"stun:stun.l.google.com:19302"}, // TODO: check radio station ice server
			},
		},
	}, wsConn, clanId, channelId, userId)
	if err != nil {
		panic(err)
	}

	err = rtcConn.SendAudioTrack(filePath)
	if err != nil {
		panic(err)
	}
}
