package toornamentClient

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestGroups(t *testing.T) {
	var client ToornamentClient
	client.ApiKey = os.Getenv("KEY")

	groups := GetGroups(&client,"4159532293277130752",GroupScope().VIEWER,new(GroupParams),NewGroupRange(0,7))
	str, err := json.Marshal(groups)
	if err != nil {
		t.Errorf("Couldn't find anything: %v",err)
	}else{
		fmt.Println(string(str))
	}


}