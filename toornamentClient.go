package toornamentClient

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/go-resty/resty"
)
type ToornamentClient struct{
	Client http.Client
	auth Authorization
	apiKey string
}

type Authorization struct {
	AccessToken string      `json:"access_token"`
	ExpiresIn   int         `json:"expires_in"`
	TokenType   string      `json:"token_type"`
	Scope       interface{} `json:"scope"`
}

func getClient(c *ToornamentClient, clientID, clientSecret, grantType string, scope []string) (ToornamentClient, error) {

	client := resty.New()

	resp, err := client.R().
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

func getSimpleClient(c *ToornamentClient, url string, headers *map[string]string) ([]byte) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalln(err)
	}
	req.Header.Set("X-Api-Key", c.apiKey)
	for key, value := range *headers {
	req.Header.Set(key, value)
	}
	resp, err := c.Client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
		var body []byte
		body, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}
		defer resp.Body.Close()

	return body
}

