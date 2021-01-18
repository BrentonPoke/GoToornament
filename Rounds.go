package toornamentClient

import (
	"encoding/json"
	"log"
	"strconv"
	"strings"

	"github.com/go-resty/resty"
)

type RoundParams struct {
	StageIds []string `json:"stage_ids"`
	StageNumbers []int `json:"stage_numbers"`
	GroupIds []string `json:"group_ids"`
	GroupNumbers []int `json:"group_numbers"`
}

func RoundScope() *apiScope {
	return &apiScope{VIEWER: "viewer", ORGANIZER: "organizer"}
}
func GetRounds(c *ToornamentClient, tournamentId, apiScope string, params RoundParams, roundRange *apiRange) []Round {
	c.client = resty.New()
	c.client.Header.Set("Accept", "application/json")
	c.client.Header.Set("X-Api-Key", c.ApiKey)
	if apiScope != RoundScope().VIEWER {
		c.client.Header.Set("Authorization", "Bearer "+c.auth.AccessToken)
	}
	if len(params.GroupNumbers) > 0 {
		sNums := make([]string, len(params.GroupNumbers))
		for i, x := range params.GroupNumbers {
			sNums[i] = strconv.Itoa(x)
		}
		c.client.QueryParam.Set("group_numbers", strings.Join(sNums, ","))
	}
	if len(params.StageNumbers) > 0 {
		sNums := make([]string, len(params.StageNumbers))
		for i, x := range params.StageNumbers {
			sNums[i] = strconv.Itoa(x)
		}
		c.client.QueryParam.Set("stage_numbers", strings.Join(sNums, ","))
	}
	if len(params.GroupIds) > 0 {
		c.client.QueryParam.Set("group_ids", strings.Join(params.GroupIds, ","))
	}
	if len(params.StageIds) > 0 {
		c.client.QueryParam.Set("stage_ids", strings.Join(params.StageIds, ","))
	}
	resp, err := c.client.R().Get("https://api.toornament.com/"+apiScope+"/v2/tournaments/"+tournamentId+"/rounds")

	if err != nil {
		log.Fatal(err)
	}
	body := resp.Body()
	rounds := make([]Round,1,roundRange.end-roundRange.begin+1)
	err = json.Unmarshal(body, &rounds)
	if err != nil {
		log.Fatalln(err)
	}
	return rounds

}