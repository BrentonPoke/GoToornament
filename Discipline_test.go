package toornamentClient

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestDiscipline(t *testing.T) {
	var client ToornamentClient
	client.ApiKey = os.Getenv("KEY")

	disciplines := GetDisciplines(&client,DisciplineScope().VIEWER,NewDisciplineRange(0,7))
	str, err := json.Marshal(disciplines)
	if err != nil {
		t.Errorf("Couldn't find anything: %v",err)
	}else{
		fmt.Println(string(str))
	}


}

func TestSingleDiscipline(t *testing.T) {
	var client ToornamentClient
	client.ApiKey = os.Getenv("KEY")

	discipline := GetDiscipline(&client,DisciplineScope().ORGANIZER,"hearthstone")
	str, err := json.Marshal(discipline)
	if err != nil {
		t.Errorf("Couldn't find anything: %v",err)
	}else{
		fmt.Println(string(str))
	}


}