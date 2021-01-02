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

func NewDisciplineRange(begin, end int) *DisciplineRange {
	d := DisciplineRange{begin: begin, end: end}
	d.drange = "disciplines=" + strconv.Itoa(d.begin) + "-" + strconv.Itoa(d.end)
	return &d
}

func GetDiscipline(c *ToornamentClient, disciplineRange *DisciplineRange) []Discipline {
	client := resty.New()
	resp, err := client.R().
		SetHeader("Accept", "application/json").
		SetHeader("X-Api-Key", c.ApiKey).
		SetHeader("range", disciplineRange.drange).
		Get("https://api.toornament.com/viewer/v2/disciplines")
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
