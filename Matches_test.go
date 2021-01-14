package toornamentClient

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestGetMatches(t *testing.T) {
	var client ToornamentClient
	client.ApiKey = os.Getenv("KEY")
	params := MatchParams{}
	params.GroupIds = []string{"4238100855809146882","4238100855809146883","4238100855809146884"}
	params.RoundNumbers = []int{1,2,3}
	params.StageIds = []string{"4159534983302389760"}
	params.RoundIds = []string{"4238100855876255960"}

	matches := GetMatches(&client,"4159532293277130752",MatchScope().VIEWER,params,NewMatchRange(0,13))
	str, err := json.Marshal(matches)
	if err != nil {
		t.Errorf("Couldn't find anything: %v",err)
	}else{
		fmt.Println(string(str))
	}
}