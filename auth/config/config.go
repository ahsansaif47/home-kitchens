package config

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

type Config struct {
	DBUrl     string
	Port      string
	JWTSecret string
	JWTExpMin string
}

var config *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		instance, err := loadConfig()
		if err != nil {
			log.Fatalf("Error loading .env file: %v", err)
		}

		config = instance
	})
	return config
}

func loadConfig() (*Config, error) {
	err := godotenv.Load()

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s, port=%s sslmode=%s",
		os.Getenv("POSTGRES_SERVER"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_SSLMODE"),
	)

	return &Config{
		DBUrl:     dsn,
		Port:      os.Getenv("PORT"),
		JWTSecret: os.Getenv("JWT_SECRET"),
		JWTExpMin: os.Getenv("JWT_EXPIRATION_MINUTES"),
	}, err

}
