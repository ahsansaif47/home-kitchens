package services

import (
	"context"
	"log"

	pb "github.com/ahsansaif47/home-kitchens/common/gRPC/generated/notifications"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type EmailClient struct {
	conn   *grpc.ClientConn
	client pb.EmailServiceClient
}

func NewEmailClient(addr string) (*EmailClient, error) {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return &EmailClient{
		conn:   conn,
		client: pb.NewEmailServiceClient(conn),
	}, err
}

func (c *EmailClient) Close() {
	if c.conn != nil {
		c.conn.Close()
	}
}

func (c *EmailClient) SendOTPEmail(ctx context.Context, from, to, subject, message, otp string) error {
	req := &pb.SendOTPRequest{
		EmailReq: &pb.SendEmailRequest{
			To:      to,
			From:    from,
			Subject: subject,
			Message: message,
		},
		Otp: otp,
	}

	resp, err := c.client.SendOTPEmail(ctx, req)
	if err != nil {
		return err
	}

	if !resp.Success {
		log.Println("Email service failed to send email")
	}

	return nil

}
