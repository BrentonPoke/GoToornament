package toornamentClient

import (
	"encoding/json"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
)

type Status interface{}
type status struct {
	RUNNING   string
	PENDING   string
	COMPLETED string
}

type Sort struct {
	STRUCTURE      string
	SCHEDULE       string
	LATEST_RESULTS string
}

type UpdateMatchParams struct {
	ScheduledDatetime time.Time `json:"scheduled_datetime"`
	PublicNote        string    `json:"public_note"`
	PrivateNote       string    `json:"private_note"`
	Opponents         []struct {
		Number     int    `json:"number"`
		Position   int    `json:"position"`
		Result     string `json:"result"`
		Rank       int    `json:"rank"`
		Forfeit    bool   `json:"forfeit"`
		Score      int    `json:"score"`
		Properties struct {
		} `json:"properties"`
	} `json:"opponents"`
}

func MatchScope() *apiScope {
	return &apiScope{VIEWER: "viewer", ORGANIZER: "organizer"}
}
func NewSort() Sort {
	return Sort{STRUCTURE: "structures", SCHEDULE: "schedule", LATEST_RESULTS: "latest_results"}
}
func NewStatus() Status {
	return status{RUNNING: "running", PENDING: "pending", COMPLETED: "completed"}
}

func NewMatchRange(begin, end int) *apiRange {
	d := apiRange{begin: begin, end: end}
	d.drange = "matches=" + strconv.Itoa(d.begin) + "-" + strconv.Itoa(d.end)
	return &d
}

func NewMatchGamesRange(begin, end int) *apiRange {
	d := apiRange{begin: begin, end: end}
	d.drange = "games=" + strconv.Itoa(d.begin) + "-" + strconv.Itoa(d.end)
	return &d
}

type MatchParams struct {
	CustomUserIdentifiers []string   `json:"custom_user_identifiers"`
	GroupIds              []string   `json:"group_ids"`
	RoundIds              []string   `json:"round_ids"`
	RoundNumbers          []int      `json:"round_numbers"`
	ParticipantIds        []string   `json:"participant_ids"`
	TournamentIds         []string   `json:"tournament_ids"`
	StageIds              []string   `json:"stage_ids"`
	Statuses              []string   `json:"statuses"`
	IsScheduled           *bool      `json:"is_scheduled"`
	IsFeatured            *bool      `json:"is_featured"`
	ScheduledBefore       *time.Time `json:"scheduled_before"`
	ScheduledAfter        *time.Time `json:"scheduled_after"`
	Sort                  string     `json:"sort"`
}

type ReportParams struct {
	TournamentIds         []string `json:"tournament_ids"`
	MatchIds              []string `json:"match_ids"`
	ParticipantIds        []string `json:"participant_ids"`
	CustomUserIdentifiers []string `json:"custom_user_identifiers"`
	Types                 []string `json:"types"`
	IsClosed              *bool    `json:"is_closed"`
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
	if params.Sort != "" {
		c.client.QueryParam.Set("sort", params.Sort)
	} else {
		c.client.QueryParam.Set("sort", "structure")
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

	if params.ScheduledBefore != nil {
		c.client.QueryParam.Set("scheduled_before", params.ScheduledBefore.Format("2006-01-02T15:04:05+07:00"))
	}
	if params.ScheduledAfter != nil {
		c.client.QueryParam.Set("scheduled_after", params.ScheduledAfter.Format("2006-01-02T15:04:05+07:00"))
	}
	if apiScope != MatchScope().VIEWER {
		if len(params.CustomUserIdentifiers) > 0 {
			c.client.QueryParam.Set("custom_user_identifiers", strings.Join(params.CustomUserIdentifiers, ","))
		}
		c.client.Header.Set("Authorization", "Bearer "+c.auth.AccessToken)
	}

	resp, err := c.client.R().
		Get("https://api.toornament.com/" + apiScope + "/v2/tournaments/" + tournamentId + "/matches")
	if err != nil {
		log.Fatal(err)
	}
	body := resp.Body()
	matches := make([]Match, 1, matchRange.end-matchRange.begin+1)
	err = json.Unmarshal(body, &matches)
	if err != nil {
		log.Fatalln(err)
	}
	return matches
}

func GetMatch(c *ToornamentClient, tournamentId, matchId string) Match {
	c.client = resty.New()
	c.client.Header.Set("Accept", "application/json")
	c.client.Header.Set("X-Api-Key", c.ApiKey)
	resp, err := c.client.R().Get("https://api.toornament.com/viewer/v2/tournaments/" + tournamentId + "/matches/" + matchId)

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

func GetDisciplineMatches(c *ToornamentClient, disciplineId string, params MatchParams, matchRange *apiRange) []Match {
	c.client = resty.New()
	c.client.Header.Set("Accept", "application/json")
	c.client.Header.Set("X-Api-Key", c.ApiKey)
	c.client.Header.Set("range", matchRange.drange)

	if len(params.Statuses) > 0 {
		c.client.QueryParam.Set("statuses", strings.Join(params.Statuses, ","))
	}

	if len(params.ParticipantIds) > 0 {
		c.client.QueryParam.Set("participant_ids", strings.Join(params.ParticipantIds, ","))
	}
	if len(params.ParticipantIds) > 0 {
		c.client.QueryParam.Set("tournament_ids", strings.Join(params.TournamentIds, ","))
	}
	if params.ScheduledBefore != nil {
		c.client.QueryParam.Set("scheduled_before", params.ScheduledBefore.Format("2006-01-02T15:04:05+07:00"))
	}
	if params.ScheduledAfter != nil {
		c.client.QueryParam.Set("scheduled_after", params.ScheduledAfter.Format("2006-01-02T15:04:05+07:00"))
	}
	if params.Sort != "" {
		c.client.QueryParam.Set("sort", params.Sort)
	} else {
		c.client.QueryParam.Set("sort", "structure")
	}
	if params.IsFeatured != nil {
		input := func(i bool) string {
			switch {
			case i:
				return "1"
			default:
				return "0"
			}
		}
		c.client.QueryParam.Set("is_featured", input(*params.IsScheduled))
	}

	resp, err := c.client.R().
		Get("https://api.toornament.com/viewer/v2/disciplines/" + disciplineId + "/matches")
	if err != nil {
		log.Fatal(err)
	}
	body := resp.Body()
	matches := make([]Match, 1, matchRange.end-matchRange.begin+1)
	err = json.Unmarshal(body, &matches)
	if err != nil {
		log.Fatalln(err)
	}
	return matches
}

func UpdateMatch(c *ToornamentClient, tournamentId, matchId string, params *UpdateMatchParams) Match {
	c.client = resty.New()
	c.client.Header.Set("Accept", "application/json")
	c.client.Header.Set("X-Api-Key", c.ApiKey)
	c.client.Header.Set("Authorization", "Bearer "+c.auth.AccessToken)
	resp, err := c.client.R().SetBody(params).Patch("https://api.toornament.com/organizer/v2/tournaments/" + tournamentId + "/matches/" + matchId)
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

func GetMatchGame(c *ToornamentClient, tournamentId, apiScope, matchId string, gameNumber int) MatchGame {
	c.client = resty.New()
	c.client.Header.Set("Accept", "application/json")
	c.client.Header.Set("X-Api-Key", c.ApiKey)
	if apiScope != MatchScope().VIEWER {
		c.client.Header.Set("Authorization", "Bearer "+c.auth.AccessToken)
	}
	resp, err := c.client.R().Get("https://api.toornament.com/" + apiScope + "/v2/tournaments/" + tournamentId + "/matches/" + matchId + "/games/" + strconv.Itoa(gameNumber))
	if err != nil {
		log.Fatal(err)
	}
	body := resp.Body()
	match := new(MatchGame)
	err = json.Unmarshal(body, &match)
	if err != nil {
		log.Fatalln(err)
	}
	return *match
}

func GetMatchGames(c *ToornamentClient, tournamentId, apiScope, matchId string, gamesRange *apiRange) []MatchGame {
	c.client = resty.New()
	c.client.Header.Set("Accept", "application/json")
	c.client.Header.Set("X-Api-Key", c.ApiKey)
	if apiScope != MatchScope().VIEWER {
		c.client.Header.Set("Authorization", "Bearer "+c.auth.AccessToken)
	}
	resp, err := c.client.R().Get("https://api.toornament.com/" + apiScope + "/v2/tournaments/" + tournamentId + "/matches/" + matchId + "/games")
	if err != nil {
		log.Fatal(err)
	}
	body := resp.Body()
	games := make([]MatchGame, 1, gamesRange.end-gamesRange.begin+1)
	err = json.Unmarshal(body, &games)
	if err != nil {
		log.Fatalln(err)
	}
	return games
}

func UpdateMatchGame(c *ToornamentClient, tournamentId, matchId string, gameNumber int, params *MatchGame) MatchGame {
	c.client = resty.New()
	c.client.Header.Set("Accept", "application/json")
	c.client.Header.Set("X-Api-Key", c.ApiKey)
	c.client.Header.Set("Authorization", "Bearer "+c.auth.AccessToken)
	resp, err := c.client.R().SetBody(params).Patch("https://api.toornament.com/organizer/v2/tournaments/" + tournamentId + "/matches/" + matchId + "/games/" + strconv.Itoa(gameNumber))
	if err != nil {
		log.Fatal(err)
	}
	body := resp.Body()
	match := new(MatchGame)
	err = json.Unmarshal(body, &match)
	if err != nil {
		log.Fatalln(err)
	}
	return *match
}

func GetMatchReports(c *ToornamentClient, reportsRange *apiRange, params *ReportParams) []MatchReport {
	c.client = resty.New()
	c.client.Header.Set("Accept", "application/json")
	c.client.Header.Set("X-Api-Key", c.ApiKey)
	c.client.Header.Set("range", reportsRange.drange)
	c.client.Header.Set("Authorization", "Bearer "+c.auth.AccessToken)
	resp, err := c.client.R().Get("https://api.toornament.com/organizer/v2/reports")
	if err != nil {
		log.Fatal(err)
	}

	if len(params.MatchIds) > 0 {
		c.client.QueryParam.Set("match_ids", strings.Join(params.MatchIds, ","))
	}

	if len(params.TournamentIds) > 0 {
		c.client.QueryParam.Set("tournament_ids", strings.Join(params.TournamentIds, ","))
	}

	if len(params.Types) > 0 {
		c.client.QueryParam.Set("types", strings.Join(params.Types, ","))
	}

	if len(params.ParticipantIds) > 0 {
		c.client.QueryParam.Set("participant_ids", strings.Join(params.ParticipantIds, ","))
	}

	if params.IsClosed != nil {
		input := func(i bool) string {
			switch {
			case i:
				return "1"
			default:
				return "0"
			}
		}
		c.client.QueryParam.Set("is_closed", input(*params.IsClosed))
	}
	body := resp.Body()
	reports := make([]MatchReport, 1, reportsRange.end-reportsRange.begin+1)
	err = json.Unmarshal(body, &reports)
	if err != nil {
		log.Fatalln(err)
	}

	return reports
}
