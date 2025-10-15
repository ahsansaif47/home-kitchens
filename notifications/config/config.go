package config

import (
	"log"
	"os"
	"path/filepath"
	"sync"

	"github.com/joho/godotenv"
)

type MailConfig struct {
	SmtpHost string
	SmtpPort string
	Email    string
	Password string
}

var config MailConfig
var once sync.Once

func GetConfig() MailConfig {
	once.Do(func() {
		instance, err := loadConfig()
		if err != nil {
			log.Fatalf("error loading config: %v", err)
		}

		config = instance
	})

	return config
}

func loadConfig() (MailConfig, error) {
	err := godotenv.Load(filepath.Join("..", "..", ".env"))

	return MailConfig{
		SmtpHost: os.Getenv("SMTP_HOST"),
		SmtpPort: os.Getenv("SMTP_PORT"),
		Email:    os.Getenv("EMAIL"),
		Password: os.Getenv("PASSWORD"),
	}, err

}
