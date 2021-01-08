package toornamentClient

import (
	"encoding/json"
	"log"
	"strconv"
	"strings"

	"github.com/go-resty/resty"
)

func RankingScope() *apiScope {
	return &apiScope{VIEWER: "viewer", ORGANIZER: "organizer"}
}

func NewRankingRange(begin, end int) *apiRange {
	d := apiRange{begin: begin, end: end}
	d.drange = "items=" + strconv.Itoa(d.begin) + "-" + strconv.Itoa(d.end)
	return &d
}

type RankingParams struct {
	CustomUserIdentifiers []string `json:"custom_user_identifiers"`
	GroupIds              []string `json:"group_ids"`
	GroupNumbers          []int `json:"group_numbers"`
	ParticipantIds        []string `json:"participant_ids"`
}

func GetRankings(c *ToornamentClient, tournamentId, stageId, apiScope string, params RankingParams, itemRange *apiRange) []Ranking {
	client := resty.New()
	client.Header.Set("Accept", "application/json")
	client.Header.Set("X-Api-Key", c.ApiKey)
	client.Header.Set("range", itemRange.drange)

	if len(params.GroupNumbers) > 0 {
		sNums := make([]string, len(params.GroupNumbers))
		for i, x := range params.GroupNumbers {
			sNums[i] = strconv.Itoa(x)
		}
		client.QueryParam.Set("group_numbers", strings.Join(sNums, ","))
	}
	if len(params.GroupIds) > 0 {
		client.QueryParam.Set("group_ids", strings.Join(params.GroupIds, ","))
	}
	if apiScope != "viewer"{
		if len(params.CustomUserIdentifiers) > 0{
			client.QueryParam.Set("custom_user_identifiers",strings.Join(params.CustomUserIdentifiers, ","))
		}

		if len(params.ParticipantIds) > 0 {
			client.QueryParam.Set("participant_ids", strings.Join(params.ParticipantIds, ","))
		}
		client.Header.Set("Authorization","Bearer "+ c.auth.AccessToken)
	}

	resp, err := client.R().
		Get("https://api.toornament.com/"+apiScope+"/v2/tournaments/"+tournamentId+"/stages/"+stageId+"/ranking-items")
	if err != nil {
		log.Fatal(err)
	}
	body := resp.Body()
	rankings := make([]Ranking,1)
	err = json.Unmarshal(body, &rankings)
	if err != nil {
		log.Fatalln(err)
	}
	return rankings
}