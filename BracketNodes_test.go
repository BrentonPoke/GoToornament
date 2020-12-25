package toornamentClient

import (
	"fmt"
	"os"
	"testing"
)

func TestBracketClient(t *testing.T) {
	var client ToornamentClient
	client.apiKey = os.Getenv("KEY")
	headers := map[string]string{"range":"nodes=0-6"}

	bracketNodes := getBracketNodes(&client,"4182678401789968384","4182681695304105984",&headers, new(BracketNodeParams))

	fmt.Println(bracketNodes)
}
