package main

import (
	"fmt"
	mezonsdk "mezon-sdk"
	"mezon-sdk/mezon-protobuf/mezon/v2/common/api"
	"mezon-sdk/mezon-protobuf/mezon/v2/common/rtapi"
	"time"
)

func main() {
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
				ClanId:           "1827955317304987648",
				ChannelId:        "1827955317309181955",
				Mode:             2,
				Content:          "{\"t\":\"Test test test\"}",
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
