package services

import (
	"context"
	"log"

	pb "github.com/ahsansaif47/home-kitchens/common/gRPC/generated"
	asyncserver "github.com/ahsansaif47/home-kitchens/notifications/async_server"
	"github.com/ahsansaif47/home-kitchens/notifications/async_server/workflows"
	"go.temporal.io/sdk/client"
)

// type EmailServiceServer interface {
// 	SendOTPEmail(context.Context, *SendEmailRequest) (*SendEmailResponse, error)
// }

type EmailServiceServer struct {
	pb.UnimplementedEmailServiceServer
	srv pb.EmailServiceServer
}

func (s *EmailServiceServer) SendOTPEmail(ctx context.Context, req *pb.SendEmailRequest) (*pb.SendEmailResponse, error) {

	options := client.StartWorkflowOptions{
		TaskQueue: "otp-queue",
	}

	we, err := asyncserver.TClient.ExecuteWorkflow(context.Background(), options, workflows.SendOTPWorkflow, req.To, req.Otp)
	if err != nil {
		log.Fatalf("unable to start workflow: %v", err)
	}

	log.Printf("Started workflow with ID: %s, RunID: %s", we.GetID(), we.GetRunID())

	return &pb.SendEmailResponse{
		Success: true,
	}, nil
}
