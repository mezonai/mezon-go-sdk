package mezonsdk

import (
	"context"
	"crypto/tls"
	"encoding/base64"
	"fmt"
	swagger "mezon-sdk/mezon-api"
	"mezon-sdk/utils"
	"net/http"
	"time"
)

type Client struct {
	api *swagger.APIClient
}

func NewClientApi(c *Config) (*swagger.MezonApiService, error) {
	token, err := getAuthenticate(c)
	if err != nil {
		return nil, err
	}

	cfg := getSwaggerConfig(c)
	cfg.AddDefaultHeader("Authorization", fmt.Sprintf("Bearer %s", token))
	return (&Client{
		api: swagger.NewAPIClient(cfg),
	}).api.MezonApi, nil
}

func getAuthenticate(c *Config) (string, error) {
	cfg := getSwaggerConfig(c)

	cfg.AddDefaultHeader("Authorization", base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("Basic %s:", c.ApiKey))))
	token, _, err := swagger.NewAPIClient(cfg).MezonApi.MezonAuthenticate(context.Background(), swagger.ApiAuthenticateRequest{
		Account: &swagger.ApiAccountApp{
			Token: c.ApiKey,
		},
	})
	if err != nil {
		return "", err
	}
	return token.Token, err
}

func getSwaggerConfig(c *Config) *swagger.Configuration {
	cfg := swagger.NewConfiguration()
	cfg.BasePath = utils.GetBasePath("http", c.BasePath, c.UseSSL)
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

	return cfg
}
