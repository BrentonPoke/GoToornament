package toornamentClient

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
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
	var sb strings.Builder
	sb.WriteString("https://api.toornament.com/oauth/v2/token")
	sb.WriteString("?grant_type="+grantType)
	sb.WriteString("&client_secret="+clientSecret)
	sb.WriteString("&client_id="+clientID)

	if scope != nil {sb.WriteString("&scope="+strings.Join(scope, ","))}

	req, err := http.NewRequest("GET", sb.String(), nil)
	if err != nil {
		log.Fatalln(err)
	}
	req.Header.Set("X-Api-Key", c.apiKey)
	resp, err := c.Client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}else{
		var body []byte
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}
		defer resp.Body.Close()
		err = json.Unmarshal(body, &c.auth)
	}
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

