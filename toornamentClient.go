package toornamentClient

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)
type toornamentClient struct{
	Client http.Client
	auth Authorization
}

type Authorization struct {
	AccessToken string      `json:"access_token"`
	ExpiresIn   int         `json:"expires_in"`
	TokenType   string      `json:"token_type"`
	Scope       interface{} `json:"scope"`
}

func getClient(c *toornamentClient, clientID, clientSecret, grantType, apiKey string, scope []string) (toornamentClient, error) {
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
	req.Header.Set("X-Api-Key", apiKey)
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
