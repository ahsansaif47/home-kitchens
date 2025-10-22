package asyncserver

import (
	"context"
	"log"

	"go.temporal.io/sdk/client"
)

var TClient client.Client

func SetupWorkers() error {
	SetupOTPWorker()

	return nil
}

func NewAsyncServer() (client.Client, error) {
	client, err := client.DialContext(context.Background(), client.Options{
		HostPort:  "localhost:7233",
		Namespace: "Notifications",
	})
	if err != nil {
		log.Fatalf("unable to create temporal client: %v", err)
	}

	SetupWorkers()
	return client, err
}
