package mezonsdk

import (
	"crypto/tls"
	"encoding/base64"
	"fmt"
	swagger "mezon-sdk/mezon-api"
	"net/http"
	"time"
)

var (
	clientBasicAuth  *Client
	clientBearerAuth *Client
)

type Config struct {
	BasePath     string   `json:"base_path"`
	Token        string   `json:"token"`
	Timeout      int      `json:"timeout"`
	InsecureSkip bool     `json:"insecure"`
	Session      *Session `json:"session"`
}

type Client struct {
	api *swagger.APIClient
}

func newClient(c *Config, session bool) *Client {
	cfg := swagger.NewConfiguration()
	cfg.BasePath = c.BasePath
	if c.Timeout == 0 {
		c.Timeout = 15
	}

	if c.InsecureSkip {
		cfg.HTTPClient = &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true,
				},
			},
		}
	} else {
		cfg.HTTPClient = http.DefaultClient
	}
	cfg.HTTPClient.Timeout = time.Duration(c.Timeout) * time.Second

	if session {
		// TODO: go ticker check session expire and renew token
		cfg.AddDefaultHeader("Authorization", fmt.Sprintf("Bearer %s", c.Token))
	} else {
		cfg.AddDefaultHeader("Authorization", base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("Basic %s:", c.Token))))
	}

	return &Client{
		api: swagger.NewAPIClient(cfg),
	}
}

func GetClientBasicAuth(cfg *Config) *swagger.MezonApiService {
	if clientBasicAuth == nil {
		clientBasicAuth = newClient(cfg, false)
	}
	return clientBasicAuth.api.MezonApi
}

func GetClientBearerAuth(cfg *Config) *swagger.MezonApiService {
	if clientBearerAuth == nil {
		clientBearerAuth = newClient(cfg, true)
	}
	return clientBearerAuth.api.MezonApi
}
