/*
 * Mezon API v2
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2.0
 * Contact: hello@heroiclabs.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type ApiGiveCoffeeEvent struct {
	SenderId string `json:"senderId,omitempty"`
	ReceiverId string `json:"receiverId,omitempty"`
	TokenCount int32 `json:"tokenCount,omitempty"`
	MessageRefId string `json:"messageRefId,omitempty"`
	ChannelId string `json:"channelId,omitempty"`
	ClanId string `json:"clanId,omitempty"`
}
