package toornamentClient

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestParticipant(t *testing.T) {
	var client ToornamentClient
	client.ApiKey = os.Getenv("KEY")

	disciplines := GetParticipants(&client, "4159532293277130752", ParticipantScope().PARTICIPANT, NewParticipantRange(0, 7))
	str, err := json.Marshal(disciplines)
	if err != nil {
		t.Errorf("Couldn't find anything: %v", err)
	} else {
		fmt.Println(string(str))
	}

}

func TestSingleParticipant(t *testing.T) {
	var client ToornamentClient
	client.ApiKey = os.Getenv("KEY")

	discipline := GetParticipant(&client, "4159532293277130752", ParticipantScope().PARTICIPANT, "4238002239210135552")
	str, err := json.Marshal(discipline)
	if err != nil {
		t.Errorf("Couldn't find anything: %v", err)
	} else {
		fmt.Println(string(str))
	}
}
