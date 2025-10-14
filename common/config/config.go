package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"

	"github.com/joho/godotenv"
)

type Config struct {
	DBUrl    string
	RedisUrl string
	Port     string
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
	err := godotenv.Load(filepath.Join("..", "..", ".env"))

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		os.Getenv("POSTGRES_SERVER"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_SSLMODE"),
	)

	redisUrl := fmt.Sprintf(
		"redis://:%s@%s:%s/%s",
		os.Getenv("REDIS_PASSWORD"),
		os.Getenv("REDIS_HOST"),
		os.Getenv("REDIS_PORT"),
		os.Getenv("REDIS_DATABASE"),
	)

	return &Config{
		DBUrl:    dsn,
		RedisUrl: redisUrl,
		Port:     os.Getenv("PORT"),
	}, err

}
