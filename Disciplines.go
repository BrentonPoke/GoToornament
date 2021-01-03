package toornamentClient

import (
	"encoding/json"
	"log"
	"strconv"

	"github.com/go-resty/resty"
)

type DisciplineRange struct {
	begin, end int
	drange     string
}

type disciplineScope struct {
	VIEWER string
	PARTICIPANT string
	ORGANIZER string
}

func DisciplineScope() *disciplineScope {
	return &disciplineScope{VIEWER: "viewer", PARTICIPANT: "participant", ORGANIZER: "organizer"}
}

func NewDisciplineRange(begin, end int) *DisciplineRange {
	d := DisciplineRange{begin: begin, end: end}
	d.drange = "disciplines=" + strconv.Itoa(d.begin) + "-" + strconv.Itoa(d.end)
	return &d
}

func GetDisciplines(c *ToornamentClient, apiScope string, disciplineRange *DisciplineRange) []Discipline {
	client := resty.New()
	resp, err := client.R().
		SetHeader("Accept", "application/json").
		SetHeader("X-Api-Key", c.ApiKey).
		SetHeader("range", disciplineRange.drange).
		Get("https://api.toornament.com/"+apiScope+"/v2/disciplines")
	if err != nil {
		log.Fatal(err)
	}
	body := resp.Body()
	discipline := make([]Discipline,1)
	err = json.Unmarshal(body, &discipline)
	if err != nil {
		log.Fatalln(err)
	}
	return discipline
}

func GetDiscipline(c *ToornamentClient, disciplineScope, id string) Discipline {
	client := resty.New()
	resp, err := client.R().
		SetHeader("Accept", "application/json").
		SetHeader("X-Api-Key", c.ApiKey).
		Get("https://api.toornament.com/"+disciplineScope+"/v2/disciplines/"+id)
	if err != nil {
		log.Fatal(err)
	}
	body := resp.Body()
	discipline := new(Discipline)
	err = json.Unmarshal(body, &discipline)
	if err != nil {
		log.Fatalln(err)
	}
	return *discipline
}