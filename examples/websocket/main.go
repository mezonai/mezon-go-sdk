package main

import (
	"fmt"
	mezonsdk "mezon-sdk"
	"mezon-sdk/mezon-protobuf/mezon/v2/common/api"
	"mezon-sdk/mezon-protobuf/mezon/v2/common/rtapi"
	"time"
)

func main() {
	conn, err := mezonsdk.GetWSConnection("wss://dev-mezon.nccsoft.vn:7305",
		"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.xxx.l_jOgFVmHZguGjAWG1h773auvrR5HMk4iamjd_xFA7g",
		"1827987498463137792")
	if err != nil {
		panic(err)
	}

	conn.SetRecvHandler(func(e *rtapi.Envelope) error {
		switch e.Message.(type) {
		case *rtapi.Envelope_ChannelMessage:
			fmt.Printf("channel message => cid: %s, message: %+v \n", e.Cid, e.GetChannelMessage())
		case *rtapi.Envelope_Pong:
			fmt.Printf("pong => cid: %s, message: %+v \n", e.Cid, e.GetPong())
		}
		return nil
	})

	time.Sleep(1 * time.Second)

	err = conn.SendMessage(&rtapi.Envelope{
		Message: &rtapi.Envelope_ChannelMessageSend{
			ChannelMessageSend: &rtapi.ChannelMessageSend{
				ClanId:           "1827987498463137792",
				ChannelId:        "1827987498479915008",
				Mode:             2,
				Content:          "{\"t\":\"Ahihi\"}",
				Mentions:         []*api.MessageMention{},
				Attachments:      []*api.MessageAttachment{},
				AnonymousMessage: false,
				MentionEveryone:  false,
				Avatar:           "",
				IsPublic:         true,
				Code:             0,
			},
		},
	})
	if err != nil {
		panic(err)
	}

	select {}
}
