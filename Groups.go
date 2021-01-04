package toornamentClient

import (
	"encoding/json"
	"log"
	"strconv"
	"strings"

	"github.com/go-resty/resty"
)


func GroupScope() *apiScope {
	return &apiScope{VIEWER: "viewer", ORGANIZER: "organizer"}
}

func NewGroupRange(begin, end int) *apiRange {
	d := apiRange{begin: begin, end: end}
	d.drange = "groups=" + strconv.Itoa(d.begin) + "-" + strconv.Itoa(d.end)
	return &d
}

type GroupParams struct {
	StageNumbers []string `json:"stage_numbers"`
	StageIds []string `json:"stage_ids"`
}
func GetGroups(c *ToornamentClient, tournamentId, apiScope string,params *GroupParams, groupRange *apiRange) []Group {
	client := resty.New()
	client.Header.Set("Accept", "application/json")
	client.Header.Set("X-Api-Key", c.ApiKey)
	client.Header.Set("range", groupRange.drange)
	if len(params.StageNumbers) > 0 {
		client.QueryParam.Set("stage_numbers", strings.Join(params.StageNumbers,","))
	}
	if len(params.StageIds) > 0 {
		client.QueryParam.Set("stage_id", strings.Join(params.StageIds, ","))
	}
	if apiScope != "viewer"{
		client.Header.Set("Authorization","Bearer "+ c.auth.AccessToken)
	}
	resp, err := client.R().
		Get("https://api.toornament.com/"+apiScope+"/v2/tournaments/"+tournamentId+"/groups")
	if err != nil {
		log.Fatal(err)
	}
	body := resp.Body()
	groups := make([]Group,1)
	err = json.Unmarshal(body, &groups)
	if err != nil {
		log.Fatalln(err)
	}
	return groups
}