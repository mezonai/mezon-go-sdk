package main

import (
	mezonsdk "mezon-sdk"

	"github.com/pion/webrtc/v4"
)

func main() {
	conn, err := mezonsdk.GetWSConnection("wss://dev-mezon.nccsoft.vn:7305",
		"xxx.eyJ0aWQiOiI3NzEzODg3OC03N2I0LTQ5YWYtYmYyNC0yZmIxMjY4NzljNzIiLCJ1aWQiOjE4Mjc5NTgyNjIyNDc0NjA4NjQsInVzbiI6Im1pbmgubmd1eWVudmFuMSIsImV4cCI6MTczMzQ3Mzk2M30.l_jOgFVmHZguGjAWG1h773auvrR5HMk4iamjd_xFA7g",
		"1827987498463137792")
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
