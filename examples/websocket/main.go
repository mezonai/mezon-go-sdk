package main

import (
	"fmt"
	"time"

	mezonsdk "github.com/nccasia/mezon-go-sdk"
	"github.com/nccasia/mezon-go-sdk/configs"
	"github.com/nccasia/mezon-go-sdk/mezon-protobuf/mezon/v2/common/api"
	"github.com/nccasia/mezon-go-sdk/mezon-protobuf/mezon/v2/common/rtapi"
)

func main() {
	conn, err := mezonsdk.NewWSConnection(&configs.Config{
		// BasePath: "dev-mezon.nccsoft.vn:7305",
		BasePath:     "api.mezon.vn",
		ApiKey:       "7663586b61444979547175356a5a4d52",
		Timeout:      10,
		InsecureSkip: true,
		UseSSL:       true,
	}, "1827955317304987648")
	if err != nil {
		panic(err)
	}

	conn.SetOnPong(func(e *rtapi.Envelope) error {
		fmt.Printf("pong => cid: %s, message: %+v \n", e.Cid, e.GetPong())
		return nil
	})

	time.Sleep(1 * time.Second)

	conn.SetOnChannelMessage(func(e *rtapi.Envelope) error {
		fmt.Printf("messages => cid: %s, message: %+v \n", e.Cid, e.GetChannelMessage())
		return nil
	})

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
