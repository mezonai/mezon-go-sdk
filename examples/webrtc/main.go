package main

import (
	mezonsdk "mezon-sdk"

	"github.com/pion/webrtc/v4"
)

func main() {
	conn, err := mezonsdk.NewWSConnection(&mezonsdk.Config{}, "")
	if err != nil {
		panic(err)
	}

	filePath := ""
	rtcConn, err := mezonsdk.NewRTCConnection(webrtc.Configuration{}, conn, "1827987498463137792", "1827987498463137792")
	if err != nil {
		panic(err)
	}

	err = rtcConn.SendFile(filePath)
	if err != nil {
		panic(err)
	}

	select {}
}
