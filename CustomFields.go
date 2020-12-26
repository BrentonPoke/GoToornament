package toornamentClient

import (
	"encoding/json"
	"log"

	"github.com/go-resty/resty"
)

func GetCustomFields(c *ToornamentClient, tournamentId, targetType string) []CustomFields{
	c.client = resty.New()

	resp, err := c.client.R().
		SetHeader("Accept", "application/json").
		SetHeader("X-Api-Key", c.ApiKey).
		SetQueryParam("target_type",targetType).
		Get("https://api.toornament.com/viewer/v2/tournaments/"+tournamentId+"/custom-fields")
	if err != nil {
		log.Fatal(err)
	}
	body := resp.Body()
	customFields := make([]CustomFields,1)
	err = json.Unmarshal(body,&customFields)

	if err != nil {
		log.Fatal(err)
	}
	return customFields
}