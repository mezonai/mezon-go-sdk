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

type ApiMessageAttachment struct {
	Filename string `json:"filename,omitempty"`
	Size string `json:"size,omitempty"`
	Url string `json:"url,omitempty"`
	Filetype string `json:"filetype,omitempty"`
	Width int32 `json:"width,omitempty"`
	Height int32 `json:"height,omitempty"`
}
