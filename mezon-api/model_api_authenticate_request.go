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

// Authenticate against the server with a device ID.
type ApiAuthenticateRequest struct {
	// The App account details.
	Account *ApiAccountApp `json:"account,omitempty"`
}
