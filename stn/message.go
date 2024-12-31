package stn

import "encoding/json"

type WsMsg struct {
	Key       string
	ClanId    string
	ChannelId string
	UserId    string
	Username  string
	Value     json.RawMessage
}
