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

type ApiOnboardingContent struct {
	GuideType int32 `json:"guideType,omitempty"`
	TaskType int32 `json:"taskType,omitempty"`
	ChannelId string `json:"channelId,omitempty"`
	Title string `json:"title,omitempty"`
	Content string `json:"content,omitempty"`
	ImageUrl string `json:"imageUrl,omitempty"`
	Answers []ApiOnboardingAnswer `json:"answers,omitempty"`
}
