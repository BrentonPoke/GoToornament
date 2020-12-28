package toornamentClient

import (
	"fmt"
	"os"
	"testing"
)

func TestBracketClient(t *testing.T) {
	var client ToornamentClient
	client.ApiKey = os.Getenv("KEY")
	fmt.Println(client.ApiKey)
	headers := map[string]string{"range":"nodes=0-6"}

	bracketNodes := GetBracketNodes(&client,"4182678401789968384","4182681695304105984",headers, new(BracketNodeParams))

	fmt.Println(bracketNodes)
	if bracketNodes == nil {
		t.Error("Couldn't find anything")
	}
}
