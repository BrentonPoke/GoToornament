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

	disciplines := GetDiscipline(&client,NewDisciplineRange(0,7))
	str, err := json.Marshal(disciplines)
	if err != nil {
		t.Errorf("Couldn't find anything: %v",err)
	}else{
		fmt.Println(string(str))
	}


}