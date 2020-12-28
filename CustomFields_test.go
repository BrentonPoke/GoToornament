package toornamentClient

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestGetCustomFields(t *testing.T) {
	var client ToornamentClient
	client.ApiKey = os.Getenv("KEY")

	customFields := GetCustomFields(&client,"4168798370927648768","team")

	str, err := json.Marshal(customFields)
	if err == nil {
		fmt.Println(string(str))
	}else {
		t.Errorf("Couldn't find anything: %v",err)
	}
}