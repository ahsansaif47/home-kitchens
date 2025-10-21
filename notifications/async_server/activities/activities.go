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

func SendOTPActivity(ctx context.Context, recipient, subject, message, otp string) error {
	cfg := config.GetConfig()

	sender := SenderData{
		Email:    cfg.Email,
		Password: cfg.Password,
	}

	addr := fmt.Sprintf("%s:%s", cfg.SmtpHost, cfg.SmtpPort)

	msg := fmt.Sprintf(
		"From: %s\r\n"+
			"To: %s\r\n"+
			"Subject: %s\r\n"+
			"MIME-Version: 1.0\r\n"+
			"Content-Type: text/plain; charset=\"utf-8\"\r\n"+
			"\r\n"+
			"Your OTP is %s\r\n",
		sender.Email, recipient, subject, otp,
	)

	auth := smtp.PlainAuth("", sender.Email, sender.Password, cfg.SmtpHost)
	err := smtp.SendMail(addr, auth, sender.Email, []string{recipient}, []byte(msg))
	if err != nil {
		log.Println("Error sending OTP: %w", err)
		return err
	}

	log.Printf("✅ OTP email sent successfully to %s", recipient)
	return nil
}
