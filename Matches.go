package toornamentClient

import (
	"encoding/json"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/go-resty/resty"
)

type Status interface{}
type status struct {
	RUNNING   string
	PENDING   string
	COMPLETED string
}
type Sort interface{}
type sort struct {
	STRUCTURE      string
	SCHEDULE       string
	LATEST_RESULTS string
}
func MatchScope() *apiScope {
	return &apiScope{VIEWER: "viewer", ORGANIZER: "organizer"}
}
func NewSort() Sort {
	return sort{STRUCTURE: "structures", SCHEDULE: "schedule", LATEST_RESULTS: "latest_results"}
}
func NewStatus() Status {
	return status{RUNNING: "running", PENDING: "pending", COMPLETED: "completed"}
}

func NewMatchRange(begin, end int) *apiRange {
	d := apiRange{begin: begin, end: end}
	d.drange = "matches=" + strconv.Itoa(d.begin) + "-" + strconv.Itoa(d.end)
	return &d
}

type MatchParams struct {
	CustomUserIdentifiers []string  `json:"custom_user_identifiers"`
	GroupIds              []string  `json:"group_ids"`
	RoundIds              []string  `json:"round_ids"`
	RoundNumbers              []int  `json:"round_numbers"`
	ParticipantIds        []string  `json:"participant_ids"`
	StageIds              []string  `json:"stage_ids"`
	Statuses              []string  `json:"statuses"`
	IsScheduled           *bool      `json:"is_scheduled"`
	ScheduledBefore       *time.Time `json:"scheduled_before"`
	ScheduledAfter        *time.Time `json:"scheduled_after"`
	Sort                  Sort      `json:"sort"`
}

func GetMatches(c *ToornamentClient, tournamentId, apiScope string, params MatchParams, matchRange *apiRange) []Match {
	c.client = resty.New()
	c.client.Header.Set("Accept", "application/json")
	c.client.Header.Set("X-Api-Key", c.ApiKey)
	c.client.Header.Set("range", matchRange.drange)

	if len(params.GroupIds) > 0 {
		c.client.QueryParam.Set("group_ids", strings.Join(params.GroupIds, ","))
	}
	if len(params.StageIds) > 0 {
		c.client.QueryParam.Set("stage_ids", strings.Join(params.StageIds, ","))
	}

	if len(params.Statuses) > 0 {
		c.client.QueryParam.Set("statuses", strings.Join(params.Statuses, ","))
	}

	if len(params.ParticipantIds) > 0 {
		c.client.QueryParam.Set("participant_ids", strings.Join(params.ParticipantIds, ","))
	}

	if len(params.RoundIds) > 0 {
		c.client.QueryParam.Set("round_ids", strings.Join(params.RoundIds, ","))
	}

	if params.ScheduledBefore != nil{
		c.client.QueryParam.Set("scheduled_before",params.ScheduledBefore.Format("2006-01-02T15:04:05+07:00"))
	}
	if params.IsScheduled != nil {
		input := func(i bool) string {
			switch {
			case i:
				return "1"
			default:
				return "0"
			}
		}
		c.client.QueryParam.Set("is_scheduled", input(*params.IsScheduled))
	}

	if params.ScheduledBefore != nil{
		c.client.QueryParam.Set("scheduled_before",params.ScheduledBefore.Format("2006-01-02T15:04:05+07:00"))
	}
	if params.ScheduledAfter != nil{
		c.client.QueryParam.Set("scheduled_after",params.ScheduledAfter.Format("2006-01-02T15:04:05+07:00"))
	}
	if params.IsScheduled != nil {
		input := func(i bool) string {
			switch {
			case i:
				return "1"
			default:
				return "0"
			}
		}
		c.client.QueryParam.Set("is_scheduled", input(*params.IsScheduled))
	}
	if apiScope != "viewer"{
		if len(params.CustomUserIdentifiers) > 0{
			c.client.QueryParam.Set("custom_user_identifiers",strings.Join(params.CustomUserIdentifiers, ","))
		}
		c.client.Header.Set("Authorization","Bearer "+ c.auth.AccessToken)
	}

	resp, err := c.client.R().
		Get("https://api.toornament.com/"+apiScope+"/v2/tournaments/"+tournamentId+"/matches")
	if err != nil {
		log.Fatal(err)
	}
	body := resp.Body()
	matches := make([]Match,1,matchRange.end-matchRange.begin+1)
	err = json.Unmarshal(body, &matches)
	if err != nil {
		log.Fatalln(err)
	}
	return matches
}

func GetMatch(c *ToornamentClient, tournamentId, matchId string) Match{
	c.client = resty.New()
	c.client.Header.Set("Accept", "application/json")
	c.client.Header.Set("X-Api-Key", c.ApiKey)
	resp, err := c.client.R().Get("https://api.toornament.com/viewer/v2/tournaments/"+tournamentId+"/matches/"+matchId)

	if err != nil {
		log.Fatal(err)
	}
	body := resp.Body()
	match := new(Match)
	err = json.Unmarshal(body, &match)
	if err != nil {
		log.Fatalln(err)
	}
	return *match
}