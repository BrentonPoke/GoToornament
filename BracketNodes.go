package toornamentClient

import (
	"encoding/json"
	"log"
	"net/url"
	"strconv"
	"strings"
	"time"

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

func getBracketNodes(c *ToornamentClient, tournamentId string, stageId string, headers *map[string]string, params *BracketNodeParams) []BracketNode {
	var sb strings.Builder
	sb.WriteString("https://api.toornament.com/viewer/v2/tournaments/")
	sb.WriteString(tournamentId)
	sb.WriteString("/stages/" + stageId + "/bracket-nodes")

	u, err := url.Parse(sb.String())
	if err != nil {
		log.Fatal(err)
	}
	if params.GroupNumbers != nil {
		u.Query().Set("group_numbers", strings.Join(params.GroupNumbers, ","))
	}
	if params.GroupIDs != nil {
		u.Query().Set("group_numbers", strings.Join(params.GroupNumbers, ","))
	}

	if params.RoundIDs != nil {
		u.Query().Set("round_ids", strings.Join(params.RoundIDs, ","))
	}

	if params.RoundNumbers != nil {
		u.Query().Set("round_numbers", strings.Join(params.RoundNumbers, ","))
		sb.WriteString(strings.Join(params.RoundNumbers, ","))
	}

	if params.MaxDepth != 0 {
		u.Query().Set("max_depth", strconv.Itoa(params.MaxDepth))
	}

	if params.MinDepth != 0 {
		u.Query().Set("min_depth", strconv.Itoa(params.MinDepth))
	}

	var bracketNodes = make([]BracketNode, 1)
	var nodes = getSimpleClient(c, u.String(), headers)

	err = json.Unmarshal(nodes, &bracketNodes)
	if err != nil {
		log.Fatalln(err)
	}
	return bracketNodes
}
