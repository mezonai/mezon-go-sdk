package main

import (
	"context"
	"errors"
	"fmt"
	mezonsdk "mezon-sdk"
)

func main() {
	api := mezonsdk.GetClientBearerAuth(&mezonsdk.Config{
		BasePath:     "https://dev-mezon.nccsoft.vn:7305",
		Token:        "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.xxxxx.824JrT9MxC7gS_iqDReMX4XBlWNtFcPOL5DnyN5cuuA",
		Timeout:      10,
		InsecureSkip: true,
		Session:      nil,
	})

	if api == nil {
		panic(errors.New("get http client api mezon is nil"))
	}

	data, _, err := api.MezonGetAccount(context.Background())
	if err != nil {
		panic(err)
	}

	fmt.Printf("data: %+v \n", data)
}
