package toornamentClient

import (
	"encoding/json"
	"log"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/go-resty/resty"
)

type BracketNode struct {
	ID                string     `json:"id"`
	StageID           string     `json:"stage_id"`
	GroupID           string     `json:"group_id"`
	RoundID           string     `json:"round_id"`
	Number            int        `json:"number"`
	Type              string     `json:"type"`
	Status            string     `json:"status"`
	ScheduledDatetime time.Time  `json:"scheduled_datetime"`
	PlayedAt          time.Time  `json:"played_at"`
	Depth             int        `json:"depth"`
	Branch            string     `json:"branch"`
	Opponents         []Opponent `json:"opponents"`
}

type BracketNodeParams struct {
	GroupIDs     []string `json:"group_ids"`
	GroupNumbers []string `json:"group_numbers"`
	RoundIDs     []string `json:"round_ids"`
	RoundNumbers []string `json:"round_numbers"`
	MinDepth     int      `json:"min_depth"`
	MaxDepth     int      `json:"max_depth"`
}

func GetBracketNodes(c *ToornamentClient, tournamentId string, stageId string, headers map[string]string, params *BracketNodeParams) []BracketNode {
	var sb strings.Builder
	sb.WriteString("https://api.toornament.com/viewer/v2/tournaments/")
	sb.WriteString(tournamentId)
	sb.WriteString("/stages/" + stageId + "/bracket-nodes")

	u, err := url.Parse(sb.String())
	if err != nil {
		log.Fatal(err)
	}
	if len(params.GroupNumbers) > 0 {
		u.Query().Set("group_numbers", strings.Join(params.GroupNumbers, ","))
	}
	if len(params.GroupIDs) > 0 {
		u.Query().Set("group_ids", strings.Join(params.GroupNumbers, ","))
	}

	if len(params.RoundIDs) > 0 {
		u.Query().Set("round_ids", strings.Join(params.RoundIDs, ","))
	}

	if len(params.RoundNumbers) > 0  {
		u.Query().Set("round_numbers", strings.Join(params.RoundNumbers, ","))
		sb.WriteString(strings.Join(params.RoundNumbers, ","))
	}

	if params.MaxDepth != 0 {
		u.Query().Set("max_depth", strconv.Itoa(params.MaxDepth))
	}

	if params.MinDepth != 0 {
		u.Query().Set("min_depth", strconv.Itoa(params.MinDepth))
	}
	value := headers["range"]
	client := resty.New()
	resp, err := client.R().
		SetHeader("Accept", "application/json").
		SetHeader("X-Api-Key", c.ApiKey).
		SetHeader("range",value).
		Get(u.String())
	if err != nil {
		log.Fatal(err)
	}
	body := resp.Body()

	var bracketNodes = make([]BracketNode, 1)
	//var nodes = getSimpleClient(c, u.String(), headers)

	err = json.Unmarshal(body, &bracketNodes)
	if err != nil {
		log.Fatalln(err)
	}
	return bracketNodes
}
