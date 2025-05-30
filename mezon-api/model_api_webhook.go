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

type ApiWebhook struct {
	Id string `json:"id,omitempty"`
	WebhookName string `json:"webhookName,omitempty"`
	ChannelId string `json:"channelId,omitempty"`
	Active int32 `json:"active,omitempty"`
	Url string `json:"url,omitempty"`
	CreatorId string `json:"creatorId,omitempty"`
	CreateTime string `json:"createTime,omitempty"`
	UpdateTime string `json:"updateTime,omitempty"`
	Avatar string `json:"avatar,omitempty"`
}
