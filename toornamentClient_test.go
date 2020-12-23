package toornamentClient

import (
	"fmt"
	"os"
	"testing"
)

func TestClient(t *testing.T) {
	var client toornamentClient

	client, err := getClient(&client,os.Getenv("CLIENT"),
		os.Getenv("SECRET"),
		"client_credentials",
		os.Getenv("KEY"),
		[]string{"organizer:admin"})
	/*printing stuff*/
	str := fmt.Sprintf("%v",client.auth)
	fmt.Println(str)
	if err != nil {t.Errorf("Expected error, received %v", err)}
}
