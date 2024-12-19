package radiostation

import "encoding/json"

type WsMsg struct {
	Key       string
	ClanId    string
	ChannelId string
	UserId    string
	Value     json.RawMessage
}
