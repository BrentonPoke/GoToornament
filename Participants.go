package toornamentClient

import (
	"encoding/json"
	"log"
	"strconv"

	"github.com/go-resty/resty"
)

func ParticipantScope() *apiScope {
	return &apiScope{VIEWER: "viewer", PARTICIPANT: "participant", ORGANIZER: "organizer"}
}

func NewParticipantRange(begin, end int) *apiRange {
	d := apiRange{begin: begin, end: end}
	d.drange = "participants=" + strconv.Itoa(d.begin) + "-" + strconv.Itoa(d.end)
	return &d
}

func GetParticipant(c *ToornamentClient, tournamentId, apiScope, participantId string) Participant {
	client := resty.New()
	resp, err := client.R().
		SetHeader("Accept", "application/json").
		SetHeader("X-Api-Key", c.ApiKey).
		Get("https://api.toornament.com/" + apiScope + "/v2/tournaments/"+tournamentId+"/participants/"+participantId)
	if err != nil {
		log.Fatal(err)
	}
	body := resp.Body()
	participant := new(Participant)
	err = json.Unmarshal(body, &participant)
	if err != nil {
		log.Fatalln(err)
	}
	return *participant
}

func GetParticipants(c *ToornamentClient, tournamentId, apiScope string, participantRange *apiRange) []Participant {
	client := resty.New()
	resp, err := client.R().
		SetHeader("Accept", "application/json").
		SetHeader("X-Api-Key", c.ApiKey).
		SetHeader("range", participantRange.drange).
		Get("https://api.toornament.com/" + apiScope + "/v2/tournaments/"+tournamentId+"/participants")
	if err != nil {
		log.Fatal(err)
	}
	body := resp.Body()
	participant := make([]Participant,1)
	err = json.Unmarshal(body, &participant)
	if err != nil {
		log.Fatalln(err)
	}
	return participant
}