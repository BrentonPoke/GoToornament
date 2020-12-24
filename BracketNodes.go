package toornamentClient

import (
	"encoding/json"
	"log"
	"strings"
	"time"

	"Toornament2Go/models"
)
type Opponent models.Opponent
type BracketNode struct {
	ID                string      `json:"id"`
	StageID           string      `json:"stage_id"`
	GroupID           string      `json:"group_id"`
	RoundID           string      `json:"round_id"`
	Number            int         `json:"number"`
	Type              string      `json:"type"`
	Status            string      `json:"status"`
	ScheduledDatetime time.Time   `json:"scheduled_datetime"`
	PlayedAt          time.Time   `json:"played_at"`
	Depth             int         `json:"depth"`
	Branch            string      `json:"branch"`
	Opponents         []Opponent `json:"opponents"`
}
func getBracketNodes(c *ToornamentClient, tournamentId string, stageId string, headers map[string]string, brRange int ) []BracketNode {
	var sb strings.Builder
	sb.WriteString("https://api.toornament.com/viewer/v2/tournaments/")
	sb.WriteString(tournamentId)
	sb.WriteString("/stages/"+ stageId +"/bracket-nodes")

	var bracketnodes = make([]BracketNode, brRange)
	var nodes = getSimpleClient(c, sb.String(),&headers)

	err := json.Unmarshal(nodes, &bracketnodes)
	if err != nil {
		log.Fatalln(err)
	}
return bracketnodes
}