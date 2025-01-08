package mezonsdk

import (
	"context"
	"crypto/tls"
	"encoding/base64"
	"fmt"
	"net/http"
	"time"

	"github.com/nccasia/mezon-go-sdk/configs"
	"github.com/nccasia/mezon-go-sdk/constants"
	"github.com/nccasia/mezon-go-sdk/utils"

	swagger "github.com/nccasia/mezon-go-sdk/mezon-api"
)

type Client struct {
	cfg    *configs.Config
	token  string
	Api    *swagger.MezonApiService
	Socket IWSConnection
}

func NewClient(apiKey string) (*Client, error) {
	c := &configs.Config{
		ApiKey:  apiKey,
		Timeout: 15,
	}
	cfg := getSwaggerConfig(c)
	api := swagger.NewAPIClient(cfg).MezonApi
	token, err := getAuthenticate(c, api)
	if err != nil {
		return nil, err
	}

	cfg.AddDefaultHeader("Authorization", fmt.Sprintf("Bearer %s", token))
	return (&Client{
		cfg:   c,
		token: token,
		Api:   api,
	}), nil
}

func (c *Client) CreateSocket() (IWSConnection, error) {
	clanDescs, _, err := c.Api.MezonListClanDescs(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	listJoinedClan := make([]string, len(clanDescs.Clandesc))

	// for DM
	listJoinedClan = append(listJoinedClan, "0")
	for _, clan := range clanDescs.Clandesc {
		listJoinedClan = append(listJoinedClan, clan.ClanId)
	}
	socket, err := NewWSConnection(c.cfg, c.token, listJoinedClan)
	if err != nil {
		return nil, err
	}

	c.Socket = socket

	return socket, nil
}

func getAuthenticate(c *configs.Config, api *swagger.MezonApiService) (string, error) {
	cfg := getSwaggerConfig(c)

	cfg.AddDefaultHeader("Authorization", base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("Basic %s:", c.ApiKey))))
	token, _, err := api.MezonAuthenticate(context.Background(), swagger.ApiAuthenticateRequest{
		Account: &swagger.ApiAccountApp{
			Token: c.ApiKey,
		},
	})
	if err != nil {
		return "", err
	}
	return token.Token, err
}

func getSwaggerConfig(c *configs.Config) *swagger.Configuration {
	cfg := swagger.NewConfiguration()
	cfg.BasePath = utils.GetBasePath("http", constants.MznBasePath, constants.UseSSL)
	if c.Timeout == 0 {
		c.Timeout = 15
	}

	if constants.InsecureSkip {
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
