package models

import "time"

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
	Opponents         []Opponents `json:"opponents"`
}
type CustomFields struct {
	MachineName  string `json:"machine_name"`
	Label        string `json:"label"`
	TargetType   string `json:"target_type"`
	Type         string `json:"type"`
	DefaultValue string `json:"default_value"`
	Required     bool   `json:"required"`
	Public       bool   `json:"public"`
	Position     int    `json:"position"`
}
type Participant struct {
	ID           string       `json:"id"`
	Name         string       `json:"name"`
	CustomFields CustomFields `json:"custom_fields"`
}
type Opponents struct {
	Number       int         `json:"number"`
	Result       string      `json:"result"`
	Rank         int         `json:"rank"`
	Forfeit      bool        `json:"forfeit"`
	Score        int         `json:"score"`
	SourceType   string      `json:"source_type"`
	SourceNodeID string      `json:"source_node_id"`
	Participant  Participant `json:"participant"`
}

type Discipline struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Shortname  string `json:"shortname"`
	Fullname   string `json:"fullname"`
	Copyrights string `json:"copyrights"`
	PlatformsAvailable []string   `json:"platforms_available"`
	TeamSize           TeamSize   `json:"team_size"`
	Features           []Features `json:"features"`
}
type TeamSize struct {
	Min int `json:"min"`
	Max int `json:"max"`
}
type Options struct {
}
type Features struct {
	Name    string  `json:"name"`
	Type    string  `json:"type"`
	Options Options `json:"options"`
}