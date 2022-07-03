package toornamentClient

import (
	"encoding/json"
	"log"
	"strconv"
	"strings"

	"github.com/go-resty/resty/v2"
)

func RankingScope() *apiScope {
	return &apiScope{RESULT: "organizer:result"}
}

func NewRankingRange(begin, end int) *apiRange {
	d := apiRange{begin: begin, end: end}
	d.drange = "items=" + strconv.Itoa(d.begin) + "-" + strconv.Itoa(d.end)
	return &d
}

type RankingParams struct {
	CustomUserIdentifiers []string `json:"custom_user_identifiers"`
	GroupIds              []string `json:"group_ids"`
	GroupNumbers          []int    `json:"group_numbers"`
	ParticipantIds        []string `json:"participant_ids"`
}

func GetRankings(c *ToornamentClient, tournamentId, stageId, apiScope string, params RankingParams, itemRange *apiRange) []Ranking {
	c.client = resty.New()
	c.client.Header.Set("Accept", "application/json")
	c.client.Header.Set("X-Api-Key", c.ApiKey)
	c.client.Header.Set("range", itemRange.drange)

	if apiScope != RoundScope().RESULT {
		c.client.Header.Set("Authorization", "Bearer "+c.auth.AccessToken)
	}

	if len(params.GroupNumbers) > 0 {
		sNums := make([]string, len(params.GroupNumbers))
		for i, x := range params.GroupNumbers {
			sNums[i] = strconv.Itoa(x)
		}
		c.client.QueryParam.Set("group_numbers", strings.Join(sNums, ","))
	}
	if len(params.GroupIds) > 0 {
		c.client.QueryParam.Set("group_ids", strings.Join(params.GroupIds, ","))
	}
	if apiScope != "viewer" {
		if len(params.CustomUserIdentifiers) > 0 {
			c.client.QueryParam.Set("custom_user_identifiers", strings.Join(params.CustomUserIdentifiers, ","))
		}

		if len(params.ParticipantIds) > 0 {
			c.client.QueryParam.Set("participant_ids", strings.Join(params.ParticipantIds, ","))
		}
		c.client.Header.Set("Authorization", "Bearer "+c.auth.AccessToken)
	}

	resp, err := c.client.R().
		Get("https://api.toornament.com/organizer/v2/ranking-items")
	if err != nil {
		log.Fatal(err)
	}
	body := resp.Body()
	rankings := make([]Ranking, 1, itemRange.end-itemRange.begin+1)
	err = json.Unmarshal(body, &rankings)
	if err != nil {
		log.Fatalln(err)
	}
	return rankings
}
