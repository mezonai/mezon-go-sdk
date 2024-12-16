package main

import (
	"context"
	"fmt"

	mezonsdk "github.com/nccasia/mezon-go-sdk"
	swagger "github.com/nccasia/mezon-go-sdk/mezon-api"

	"github.com/antihax/optional"
)

func main() {
	api, err := mezonsdk.NewClientApi(&mezonsdk.Config{
		BasePath: "dev-mezon.nccsoft.vn:7305",
		// BasePath:     "api.mezon.vn",
		ApiKey:       "7663586b61xxxxxxxx356a5a4d52",
		Timeout:      10,
		InsecureSkip: true,
		UseSSL:       true,
	})

	if err != nil {
		panic(err)
	}

	data, _, err := api.MezonListChannelVoiceUsers(context.Background(), &swagger.MezonListChannelVoiceUsersOpts{
		ClanId:      optional.NewString("1827955317304987648"),
		ChannelType: optional.NewInt32(4),
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("data: %+v \n", data)
}
