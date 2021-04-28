package toornamentClient

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestStages(t *testing.T) {
	var client ToornamentClient
	client.ApiKey = os.Getenv("KEY")

	stages := GetStages(&client,"4159532293277130752",StageScope().VIEWER)
	str, err := json.Marshal(stages)
	if err != nil {
		t.Errorf("Couldn't find anything: %v",err)
	}else{
		fmt.Println(string(str))
	}
}