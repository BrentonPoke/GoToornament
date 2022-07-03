package toornamentClient

import (
	"encoding/json"
	"log"
	"strconv"
	"strings"

	"github.com/go-resty/resty/v2"
)

type RoundParams struct {
	StageIds     []string `json:"stage_ids"`
	StageNumbers []int    `json:"stage_numbers"`
	GroupIds     []string `json:"group_ids"`
	GroupNumbers []int    `json:"group_numbers"`
}

func RoundScope() *apiScope {
	return &apiScope{VIEWER: "organizer:viewer", RESULT: "organizer:result"}
}
func NewRoundRange(begin, end int) *apiRange {
	d := apiRange{begin: begin, end: end}
	d.drange = "rounds=" + strconv.Itoa(d.begin) + "-" + strconv.Itoa(d.end)
	return &d
}

func GetRound(c *ToornamentClient, apiScope string, roundId string) Round {
	c.client = resty.New()
	c.client.Header.Set("Accept", "application/json")
	c.client.Header.Set("X-Api-Key", c.ApiKey)
	if apiScope != RoundScope().RESULT {
		c.client.Header.Set("Authorization", "Bearer "+c.auth.AccessToken)
	}

	resp, err := c.client.R().Get("https://api.toornament.com/organizer/v2/rounds/" + roundId)
	if err != nil {
		log.Fatal(err)
	}
	body := resp.Body()
	round := new(Round)
	err = json.Unmarshal(body, round)
	if err != nil {
		log.Fatalln(err)
	}
	return *round
}

func GetRounds(c *ToornamentClient, tournamentId, apiScope string, params RoundParams, roundRange *apiRange) []Round {
	c.client = resty.New()
	c.client.Header.Set("Accept", "application/json")
	c.client.Header.Set("X-Api-Key", c.ApiKey)
	if apiScope != RoundScope().VIEWER || apiScope != RoundScope().RESULT {
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
	resp, err := c.client.R().Get("https://api.toornament.com/organizer/v2/rounds")

	if err != nil {
		log.Fatal(err)
	}
	body := resp.Body()
	rounds := make([]Round, 1, roundRange.end-roundRange.begin+1)
	err = json.Unmarshal(body, &rounds)
	if err != nil {
		log.Fatalln(err)
	}
	return rounds

}
