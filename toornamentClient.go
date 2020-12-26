package toornamentClient

import (
	"encoding/json"
	"log"
	"strings"

	"github.com/go-resty/resty"
)
type ToornamentClient struct{
	client *resty.Client
	auth Authorization
	ApiKey string
}

type Authorization struct {
	AccessToken string      `json:"access_token"`
	ExpiresIn   int         `json:"expires_in"`
	TokenType   string      `json:"token_type"`
	Scope       interface{} `json:"scope"`
}

func getClient(c *ToornamentClient, clientID, clientSecret, grantType string, scope []string) (ToornamentClient, error) {

	c.client = resty.New()

	resp, err := c.client.R().
		SetQueryParams(map[string]string{
			"grant_type": grantType,
			"client_secret": clientSecret,
			"client_id": clientID,
			"scope": strings.Join(scope, " "),
		}).
		Get("https://api.toornament.com/oauth/v2/token")

		body := resp.Body()
		if err != nil {
			log.Fatalln(err)
		}
		err = json.Unmarshal(body, &c.auth)
	return *c, err
}
