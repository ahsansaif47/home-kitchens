package main

import (
	"context"
	"log"
	"time"

	"github.com/ahsansaif47/home-kitchens/notifications/activities"
	"github.com/ahsansaif47/home-kitchens/notifications/workflows"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func main() {

	// activities.SendOTPActivity(context.Background(), "ahsansaif047@gmail.com", "123456")

	temporalClient, err := client.DialContext(context.Background(), client.Options{
		HostPort: "localhost:7233",
	})
	if err != nil {
		log.Fatalf("unable to create temporal client: %v", err)
	}

	defer temporalClient.Close()

	w := worker.New(temporalClient, "otp-queue", worker.Options{})

	w.RegisterWorkflow(workflows.SendOTPWorkflow)
	w.RegisterActivity(activities.SendOTPActivity)

	log.Println("Worked started successfully. Listening for workflows!")

	go func() {
		err = w.Run(worker.InterruptCh())
		if err != nil {
			log.Fatalf("unable to start worker: %v", err)
		}
	}()

	// Workflow start options
	options := client.StartWorkflowOptions{
		ID:        "send-otp-workflow", // unique ID per workflow instance
		TaskQueue: "otp-queue",         // must match your worker’s Task Queue
	}

	// TODO: This needs to be called by the RPC Server.
	// Here it is used for testing purpose to check OTP flow.
	we, err := temporalClient.ExecuteWorkflow(context.Background(), options, workflows.SendOTPWorkflow, "ahsansaif047@gmail.com", "123456")
	if err != nil {
		log.Fatalf("unable to start workflow: %v", err)
	}

	log.Printf("Started workflow with ID: %s, RunID: %s", we.GetID(), we.GetRunID())

	time.Sleep(10 * time.Minute)
}
