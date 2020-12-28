package toornamentClient

import (
	"fmt"
	"os"
	"testing"
)

func TestGetCustomFields(t *testing.T) {
	var client ToornamentClient
	client.ApiKey = os.Getenv("KEY")

	customFields := GetCustomFields(&client,"4168798370927648768","team")
	fmt.Println(customFields)
	if customFields == nil {
		t.Error("Couldn't find anything")
	}
}