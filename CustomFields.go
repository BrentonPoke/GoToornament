package toornamentClient

import (
	"encoding/json"
	"log"

	"github.com/go-resty/resty"
)

func GetViewerCustomFields(c *ToornamentClient, tournamentId, targetType string) []CustomFields {
	return getCustomFields(c, tournamentId, targetType, "viewer")
}
func GetParticipantCustomFields(c *ToornamentClient, tournamentId, targetType string) []CustomFields {
	return getCustomFields(c, tournamentId, targetType, "participant")
}
func GetOrganizerCustomFields(c *ToornamentClient, tournamentId, targetType string) []CustomFields {
	c.client = resty.New()
	resp, err := c.client.R().
		SetHeader("Accept", "application/json").
		SetHeader("X-Api-Key", c.ApiKey).
		SetHeader("Authorization","Bearer "+ c.auth.AccessToken).
		SetQueryParam("target_type",targetType).
		Get("https://api.toornament.com/organizer/v2/tournaments/"+tournamentId+"/custom-fields")

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
func CreateCustomField ( c *ToornamentClient, tournamentId string, customFields *CustomFields) CustomFields{
	c.client = resty.New()
	resp, err := c.client.R().
		SetHeader("Accept", "application/json").
		SetHeader("X-Api-Key", c.ApiKey).
		SetHeader("Authorization","Bearer "+ c.auth.AccessToken).
		SetBody(customFields).
		Post("https://api.toornament.com/organizer/v2/tournaments/"+tournamentId+"/custom-fields")
	if err != nil {
		log.Fatal(err)
	}
	body := resp.Body()
	err = json.Unmarshal(body,customFields) //reusing object
	if err != nil {
		log.Fatal(err)
	}
	return *customFields
}
func getCustomFields(c *ToornamentClient, tournamentId, targetType, scope string) []CustomFields{
	c.client = resty.New()

	resp, err := c.client.R().
		SetHeader("Accept", "application/json").
		SetHeader("X-Api-Key", c.ApiKey).
		SetQueryParam("target_type",targetType).
		Get("https://api.toornament.com/"+scope+"/v2/tournaments/"+tournamentId+"/custom-fields")
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