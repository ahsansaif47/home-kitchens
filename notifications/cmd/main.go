package main

import (
	"context"
	"log"
	"time"

	asyncserver "github.com/ahsansaif47/home-kitchens/notifications/async_server"
	"github.com/ahsansaif47/home-kitchens/notifications/async_server/workflows"
	"go.temporal.io/sdk/client"
)

// TODO: Register Workers, Workflows + Activities Function

func main() {

	// activities.SendOTPActivity(context.Background(), "ahsansaif047@gmail.com", "123456")

	tClient, err := asyncserver.NewAsyncServer()
	if err != nil {
		log.Fatalf("unable to start worker: %v", err)
	}
	asyncserver.TClient = tClient

	defer asyncserver.TClient.Close()

	// w := worker.New(temporalClient, "otp-queue", worker.Options{})

	// w.RegisterWorkflow(workflows.SendOTPWorkflow)
	// w.RegisterActivity(activities.SendOTPActivity)

	// log.Println("Worked started successfully. Listening for workflows!")

	// go func() {
	// 	err = w.Run(worker.InterruptCh())
	// 	if err != nil {
	// 		log.Fatalf("unable to start worker: %v", err)
	// 	}
	// }()

	asyncserver.SetupOTPWorker()

	// TODO: This needs to be called by the RPC Server.
	// Here it is used for testing purpose to check OTP flow.
	options := client.StartWorkflowOptions{
		TaskQueue: "otp-queue", // must match your worker’s Task Queue
	}

	we, err := tClient.ExecuteWorkflow(context.Background(), options, workflows.SendOTPWorkflow, "ahsansaif047@gmail.com", "123456")
	if err != nil {
		log.Fatalf("unable to start workflow: %v", err)
	}

	log.Printf("Started workflow with ID: %s, RunID: %s", we.GetID(), we.GetRunID())

	time.Sleep(10 * time.Minute)
}
