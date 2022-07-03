package toornamentClient

import (
	"encoding/json"
	"log"
	"strconv"
	"strings"

	"github.com/go-resty/resty/v2"
)

func StageScope() *apiScope {
	return &apiScope{ADMIN: "organizer:admin", RESULT: "organizer:result"}
}

func NewStagesRange(begin, end int) *apiRange {
	d := apiRange{begin: begin, end: end}
	d.drange = "stages=" + strconv.Itoa(d.begin) + "-" + strconv.Itoa(d.end)
	return &d
}

func GetStages(c *ToornamentClient, tournamentId, apiScope string) []Stage {
	client := resty.New()
	client.Header.Set("Accept", "application/json")
	client.Header.Set("X-Api-Key", c.ApiKey)
	if apiScope != StageScope().RESULT || apiScope != StageScope().ADMIN {
		client.Header.Set("Authorization", "Bearer "+c.auth.AccessToken)
	}
	resp, err := client.R().
		Get("https://api.toornament.com/organizer/v2/stages")
	if err != nil {
		log.Fatal(err)
	}
	body := resp.Body()
	stages := make([]Stage, 1)
	err = json.Unmarshal(body, &stages)
	if err != nil {
		log.Fatalln(err)
	}
	return stages
}

func GetStagesForTournaments(c *ToornamentClient, tournamentIds []string, apiScope string, stageRange *apiRange) []Stage {
	client := resty.New()
	client.Header.Set("Accept", "application/json")
	client.Header.Set("X-Api-Key", c.ApiKey)
	client.Header.Set("Authorization", "Bearer "+c.auth.AccessToken)
	client.Header.Set("range", stageRange.drange)
	client.QueryParam.Set("tournament_ids", strings.Join(tournamentIds, ","))

	resp, err := client.R().
		Get("https://api.toornament.com/organizer/v2/stages")
	if err != nil {
		log.Fatal(err)
	}
	body := resp.Body()

	stages := make([]Stage, 1, stageRange.end+stageRange.begin+1)
	err = json.Unmarshal(body, &stages)
	if err != nil {
		log.Fatalln(err)
	}
	return stages
}

func GetStage(c *ToornamentClient, tournamentId, id, apiScope string) Stage {
	client := resty.New()
	client.Header.Set("Accept", "application/json")
	client.Header.Set("X-Api-Key", c.ApiKey)
	if apiScope != "viewer" {
		client.Header.Set("Authorization", "Bearer "+c.auth.AccessToken)
	}
	resp, err := client.R().
		Get("https://api.toornament.com/organizer/v2/stages" + id)
	if err != nil {
		log.Printf("Called with scope(s): %v", c.auth.Scope)
		log.Fatal(err)
	}
	body := resp.Body()
	stage := new(Stage)
	err = json.Unmarshal(body, &stage)
	if err != nil {
		log.Fatalln(err)
	}
	return *stage
}

func GetStageOnlyByID(c *ToornamentClient, id string) Stage {
	client := resty.New()
	client.Header.Set("Accept", "application/json")
	client.Header.Set("X-Api-Key", c.ApiKey)

	resp, err := client.R().
		Get("https://api.toornament.com/organizer/v2/stages/" + id)
	if err != nil {
		log.Printf("Called with scope(s): %v", c.auth.Scope)
		log.Fatal(err)
	}
	body := resp.Body()
	stage := new(Stage)
	err = json.Unmarshal(body, &stage)
	if err != nil {
		log.Fatalln(err)
	}
	return *stage
}
