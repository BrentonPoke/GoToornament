package toornamentClient

import (
	"encoding/json"
	"log"

	"github.com/go-resty/resty/v2"
)

type CustomFieldsQuery struct {
	Label        string `json:"label"`
	DefaultValue string `json:"default_value"`
	Required     bool   `json:"required"`
	Public       bool   `json:"public"`
	Position     int    `json:"position"`
}

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
		SetHeader("Authorization", "Bearer "+c.auth.AccessToken).
		SetQueryParam("target_type", targetType).
		Get("https://api.toornament.com/organizer/v2/tournaments/" + tournamentId + "/custom-fields")

	if err != nil {
		log.Fatal(err)
	}
	body := resp.Body()
	customFields := make([]CustomFields, 1)
	err = json.Unmarshal(body, &customFields)

	if err != nil {
		log.Fatal(err)
	}
	return customFields
}
func CreateCustomField(c *ToornamentClient, tournamentId string, customFields *CustomFields) CustomFields {
	c.client = resty.New()
	resp, err := c.client.R().
		SetHeader("Accept", "application/json").
		SetHeader("X-Api-Key", c.ApiKey).
		SetHeader("Authorization", "Bearer "+c.auth.AccessToken).
		SetBody(customFields).
		Post("https://api.toornament.com/organizer/v2/tournaments/" + tournamentId + "/custom-fields")
	if err != nil {
		log.Fatal(err)
	}
	body := resp.Body()
	err = json.Unmarshal(body, customFields) //reusing object
	if err != nil {
		log.Fatal(err)
	}
	return *customFields
}
func GetCustomField(c *ToornamentClient, tournamentId, fieldId string) CustomFields {
	c.client = resty.New()
	resp, err := c.client.R().
		SetHeader("Accept", "application/json").
		SetHeader("X-Api-Key", c.ApiKey).
		SetHeader("Authorization", "Bearer "+c.auth.AccessToken).
		Get("https://api.toornament.com/organizer/v2/tournaments/" + tournamentId + "/custom-fields" + fieldId)
	if err != nil {
		log.Fatal(err)
	}

	f := new(CustomFields)
	err = json.Unmarshal(resp.Body(), f)
	if err != nil {
		log.Fatal(err)
	}
	return *f
}
func UpdateCustomField(c *ToornamentClient, tournamentId, fieldId string, query CustomFieldsQuery) CustomFields {
	c.client = resty.New()
	body, err := json.Marshal(query)
	if err != nil {
		log.Fatal(err)
	}
	resp, err := c.client.R().
		SetHeader("Accept", "application/json").
		SetHeader("X-Api-Key", c.ApiKey).
		SetHeader("Authorization", "Bearer "+c.auth.AccessToken).
		SetBody(string(body)).
		Patch("https://api.toornament.com/organizer/v2/tournaments/" + tournamentId + "/custom-fields" + fieldId)

	if err != nil {
		log.Fatal(err)
	}

	f := new(CustomFields)
	err = json.Unmarshal(resp.Body(), f)
	if err != nil {
		log.Fatal(err)
	}
	return *f
}
func DeleteCustomField(c *ToornamentClient, tournamentId, fieldId string) {
	c.client = resty.New()
	_, err := c.client.R().
		SetHeader("Accept", "application/json").
		SetHeader("X-Api-Key", c.ApiKey).
		SetHeader("Authorization", "Bearer "+c.auth.AccessToken).
		Delete("https://api.toornament.com/organizer/v2/tournaments/" + tournamentId + "/custom-fields" + fieldId)
	if err != nil {
		log.Fatal(err)
	}

}
func getCustomFields(c *ToornamentClient, tournamentId, targetType, scope string) []CustomFields {
	c.client = resty.New()

	resp, err := c.client.R().
		SetHeader("Accept", "application/json").
		SetHeader("X-Api-Key", c.ApiKey).
		SetQueryParam("target_type", targetType).
		Get("https://api.toornament.com/" + scope + "/v2/tournaments/" + tournamentId + "/custom-fields")
	if err != nil {
		log.Fatal(err)
	}
	body := resp.Body()
	customFields := make([]CustomFields, 1)
	err = json.Unmarshal(body, &customFields)

	if err != nil {
		log.Fatal(err)
	}
	return customFields
}
