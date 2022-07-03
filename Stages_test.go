package toornamentClient

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestStages(t *testing.T) {
	var c ToornamentClient
	scope := make([]string, 2)
	scope = append(scope, "organizer:result", "organizer:participant")
	client, err := GetClient(&c, os.Getenv("CLIENT"), os.Getenv("SECRET"), "client_credentials", scope)
	client.ApiKey = os.Getenv("KEY")
	//client.auth.AccessToken = "eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIsImp0aSI6ImNmM2FlNTM2ZDlkMzY2YTQ2NzcxYWQ2NWMxNTVjYTU2MTAyNWIwYTY1ZWRlOGE3MDNiNWVhZTY2MjZjNGIzMmQ4NDg3N2YyYjNhNGFhNmRmIn0.eyJhdWQiOiI1OGZmNDQwMTE0MGJhMDhlN2Y4YjQ1NjcyNjlsdHBwd240ODBnZ3cwOGdjNGdna2NjY3dzZ3NvZzRjc3NjODBjOHN3Z2t3ZzBzbyIsImp0aSI6ImNmM2FlNTM2ZDlkMzY2YTQ2NzcxYWQ2NWMxNTVjYTU2MTAyNWIwYTY1ZWRlOGE3MDNiNWVhZTY2MjZjNGIzMmQ4NDg3N2YyYjNhNGFhNmRmIiwiaWF0IjoxNjU2ODgyMTQ0LCJuYmYiOjE2NTY4ODIxNDQsImV4cCI6MTY1Njk3MjE0NCwic3ViIjoiIiwic2NvcGVzIjpbIm9yZ2FuaXplcjpyZXN1bHQiLCJvcmdhbml6ZXI6cGFydGljaXBhbnQiXX0.GIqDwvEN3hWFm6Lhpfum8iu1BJA38o8EIdCdjli6fE2tOw-6TBDUn9T1J1boyEzM7rc69nIAjfVnGKsFCfISDFZjG-QZqPTddUKOqBe5czfsoAH4sB9tqjrUWvLDe5zb6LEjzKJOGY3uT6fhltUOOErV_OAg5iUn65KkDZDvVVmY-cMRvFTWoEBR-fmvGQBBtsCc817Gs5qV-N3mMrABfQZ7O9VCsKbqbspygIiOm_oH-0jH7FwQB_EkcRDfRVTPJGWypFmr33zTk8ZD354pIwJNEa9AZaP9NidyzkGn6y1D4Z4q10ICzhmXnZQQjUpFUMu01yxH8jWO9uVEn35YWg"
	fmt.Println(client.auth)

	tournamentIds := make([]string, 1)
	tournamentIds = append(tournamentIds, "4159532293277130752")

	stages := GetStagesForTournaments(&client, tournamentIds, StageScope().RESULT, NewStagesRange(0, 5))
	str, err := json.Marshal(stages)
	if err != nil {
		t.Errorf("Couldn't find anything: %v", err)
	} else {
		fmt.Println(string(str))
	}
}
