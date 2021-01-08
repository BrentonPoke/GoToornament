package toornamentClient

import "time"

type CustomFields struct {
	MachineName  string `json:"machine_name"`
	Label        string `json:"label"`
	TargetType   string `json:"target_type"`
	Type         string `json:"type"`
	DefaultValue string `json:"default_value"`
	Required     bool   `json:"required"`
	Public       bool   `json:"public"`
	Position     int    `json:"position"`
	Twitter      string `json:"twitter"`
	Facebook     string `json:"facebook"`
	Snapchat     string `json:"snapchat"`
	Twitch       string `json:"twitch"`
	Youtube      string `json:"youtube"`
	Instagram    string `json:"instagram"`
	Vimeo        string `json:"vimeo"`

	Address             Address  `json:"address"`
	Fullname            Fullname `json:"full_name"`
	Birthdate           string   `json:"birth_date"`
	Country             string   `json:"country"`
	HSPickChoice        string   `json:"hs_pick_choice"`
	BattleNetPlayerID   string   `json:"battle_net_player_id"`
	BloodBowl2PlayerID  string   `json:"blood_bowl2_player_id"`
	ManiaplanetPlayerID string   `json:"maniaplanet_player_id"`
	OriginPlayerID      string   `json:"origin_player_id"`
	PSNPlayerID         string   `json:"psn_player_id"`
	SmitePlayerID       string   `json:"smite_player_id"`
	RiotPlayerID        string   `json:"riot_player_id"`
	SteamPlayerID       string   `json:"steam_player_id"`
	SummonerPlayerID    string   `json:"summoner_player_id"`
	UplayPlayerID       string   `json:"uplay_player_id"`
	WargamingPlayerID   string   `json:"wargaming_player_id"`
	XboxLivePlayerID    string   `json:"xbox_live_player_id"`
	Discord             string   `json:"discord"`
	Checkbox            bool     `json:"checkbox"`
	OptIn               bool     `json:"optin"`
	Website             string   `json:"website"`
	LogoSmall           string   `json:"logo_small"`
	LogoMedium          string   `json:"logo_medium"`
	LogoLarge           string   `json:"logo_large"`
	Original            string   `json:"original"`
}
type Address struct {
	Address    string `json:"address"`
	PostalCode string `json:"postal_code"`
	City       string `json:"city"`
	Country    string `json:"country"`
}

type Fullname struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type Participant struct {
	ID                   string        `json:"id"`
	Name                 string        `json:"name"`
	CustomFields         CustomFields  `json:"custom_fields"`
	Email                string        `json:"email,omitempty"`
	CustomUserIdentifier string        `json:"custom_user_identifier,omitempty"`
	CheckedIn            bool          `json:"checked_in,omitempty"`
	UserID               string        `json:"user_id,omitempty"`
	CreatedAt            time.Time     `json:"created_at,omitempty"`
	Lineup               []Participant `json:"lineup,omitempty"`
}

type Opponent struct {
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
	ID                 string     `json:"id"`
	Name               string     `json:"name"`
	Shortname          string     `json:"shortname"`
	Fullname           string     `json:"fullname"`
	Copyrights         string     `json:"copyrights"`
	PlatformsAvailable []string   `json:"platforms_available,omitempty"`
	TeamSize           TeamSize   `json:"team_size,omitempty"`
	Features           []Features `json:"features,omitempty"`
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
type Group struct {
	ID       string   `json:"id"`
	StageID  string   `json:"stage_id"`
	Number   int      `json:"number"`
	Name     string   `json:"name"`
	Closed   bool     `json:"closed"`
	Settings Settings `json:"settings"`
}
type Settings struct {
	Size int `json:"size"`
}

type apiRange struct {
	begin, end int
	drange     string
}

type apiScope struct {
	VIEWER string
	PARTICIPANT string
	ORGANIZER string
}

type Ranking struct {
	ID          string      `json:"id"`
	GroupID     string      `json:"group_id"`
	Number      int         `json:"number"`
	Position    int         `json:"position"`
	Rank        int         `json:"rank"`
	Participant Participant `json:"participant"`
	Points      int         `json:"points"`
	Properties  Properties  `json:"properties"`
}

type Properties struct {
	Wins            int `json:"wins"`
	Draws           int `json:"draws"`
	Losses          int `json:"losses"`
	Played          int `json:"played"`
	Forfeits        int `json:"forfeits"`
	ScoreFor        int `json:"score_for"`
	ScoreAgainst    int `json:"score_against"`
	ScoreDifference int `json:"score_difference"`
}