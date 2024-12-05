package client

import (
	"encoding/base64"
	"fmt"
	swagger "mezon-sdk/mezon-api"
	"time"
)

var (
	clientBasicAuth  *Client
	clientBearerAuth *Client
)

type Config struct {
	BasePath  string   `json:"base_path"`
	ServerKey string   `json:"server_key"`
	Timeout   int      `json:"timeout"`
	Session   *Session `json:"session"`
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
	cfg.HTTPClient.Timeout = time.Duration(c.Timeout) * time.Second

	if session {
		// TODO: go ticker check session expire and renew token
		// cfg.AddDefaultHeader("Authorization", "Bearer ")
	} else {
		cfg.AddDefaultHeader("Authorization", base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("Basic %s:", c.ServerKey))))
	}

	return &Client{
		api: swagger.NewAPIClient(cfg),
	}
}

func (c *Client) GetClientBasicAuth(cfg *Config) *swagger.MezonApiService {
	if clientBasicAuth == nil {
		clientBasicAuth = newClient(cfg, false)
	}
	return clientBasicAuth.api.MezonApi
}

func (c *Client) GetClientBearerAuth(cfg *Config) *swagger.MezonApiService {
	if clientBearerAuth == nil {
		clientBearerAuth = newClient(cfg, true)
	}
	return clientBearerAuth.api.MezonApi
}
