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

// Authenticate against the server with a refresh token.
type ApiSessionRefreshRequest struct {
	// Refresh token.
	Token string `json:"token,omitempty"`
	// Extra information that will be bundled in the session token.
	Vars map[string]string `json:"vars,omitempty"`
}
