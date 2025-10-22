package main

import (
	"fmt"
	"log"
	"net"
	"time"

	asyncserver "github.com/ahsansaif47/home-kitchens/notifications/async_server"
	"github.com/ahsansaif47/home-kitchens/notifications/grpc/services"
	"google.golang.org/grpc"

	pb "github.com/ahsansaif47/home-kitchens/common/gRPC/generated/notifications"
)

func startGRPC() {
	port := 50052
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %s", err.Error())
	}

	grpcServer := grpc.NewServer()
	pb.RegisterEmailServiceServer(grpcServer, &services.EmailServiceServer{})
	log.Printf("🚀 gRPC Notification Server listening on port %d", port)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err.Error())
	}
}

func main() {
	tClient, err := asyncserver.NewAsyncServer()
	if err != nil {
		log.Fatalf("unable to start worker: %v", err)
	}
	asyncserver.TClient = tClient

	defer asyncserver.TClient.Close()

	// // TODO: This needs to be called by the RPC Server.
	// // Here it is used for testing purpose to check OTP flow.
	// options := client.StartWorkflowOptions{
	// 	TaskQueue: "otp-queue", // must match your worker’s Task Queue
	// }

	// we, err := tClient.ExecuteWorkflow(context.Background(), options, workflows.SendOTPWorkflow, "ahsansaif047@gmail.com", "123456")
	// if err != nil {
	// 	log.Fatalf("unable to start workflow: %v", err)
	// }

	// log.Printf("Started workflow with ID: %s, RunID: %s", we.GetID(), we.GetRunID())

	go func() {
		startGRPC()
	}()

	time.Sleep(10 * time.Minute)
}
