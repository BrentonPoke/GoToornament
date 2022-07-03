package toornamentClient

import (
	"encoding/json"
	"fmt"

	"os"
	"testing"
)

//I do not get results for some reason. Neither this test nor Postman gets anything
func TestRankings(t *testing.T) {
	var client ToornamentClient
	client.ApiKey = os.Getenv("KEY")
	params := RankingParams{}
	params.GroupIds = []string{"4238100855809146882", "4238100855809146883", "4238100855809146884"}
	params.GroupNumbers = []int{1, 2, 3}
	rankings := GetRankings(&client, "4159532293277130752", "4159534983302389760", RankingScope().RESULT, params, NewRankingRange(0, 7))
	str, err := json.Marshal(rankings)
	if err != nil {
		t.Errorf("Couldn't find anything: %v", err)
	} else {
		fmt.Println(string(str))
	}
}
