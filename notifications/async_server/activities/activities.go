package activities

import (
	"context"
	"fmt"
	"log"
	"net/smtp"

	"github.com/ahsansaif47/home-kitchens/notifications/config"
)

type SenderData struct {
	Email    string
	Password string
}

func SendOTPActivity(ctx context.Context, recipient, otp string) error {
	cfg := config.GetConfig()

	sender := SenderData{
		Email:    cfg.Email,
		Password: cfg.Password,
	}

	addr := fmt.Sprintf("%s:%s", cfg.SmtpHost, cfg.SmtpPort)

	message := fmt.Sprintf(
		"From: %s\r\n"+
			"To: %s\r\n"+
			"Subject: OTP Verification Email\r\n"+
			"MIME-Version: 1.0\r\n"+
			"Content-Type: text/plain; charset=\"utf-8\"\r\n"+
			"\r\n"+
			"Your OTP is %s\r\n",
		sender.Email, recipient, otp,
	)

	auth := smtp.PlainAuth("", sender.Email, sender.Password, cfg.SmtpHost)
	err := smtp.SendMail(addr, auth, sender.Email, []string{recipient}, []byte(message))
	if err != nil {
		log.Println("Error sending OTP: %w", err)
		return err
	}

	log.Printf("✅ OTP email sent successfully to %s", recipient)
	return nil
}
