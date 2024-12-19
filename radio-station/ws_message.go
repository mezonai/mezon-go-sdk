package radiostation

import "encoding/json"

type WsMsg struct {
	Key         string
	ClanId      string
	ChannelId   string
	UserId      string
	DisplayName string
	Value       json.RawMessage
}
