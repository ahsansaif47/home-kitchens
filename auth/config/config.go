package config

import (
	"log"
	"os"
	"path/filepath"
	"sync"

	"github.com/ahsansaif47/home-kitchens/common/config"
	"github.com/joho/godotenv"
)

type AuthConfig struct {
	GlobalCfg config.Config
	JWTSecret string
	JWTExpMin string
}

var conf AuthConfig
var once sync.Once

func GetConfig() AuthConfig {
	once.Do(func() {
		instance, err := loadConfig()
		if err != nil {
			log.Fatalf("Error loading .env file: %v", err)
		}

		conf = instance
	})
	return conf
}

func loadConfig() (AuthConfig, error) {
	err := godotenv.Load(filepath.Join("..", "..", ".env"))

	globalConf := config.GetConfig()

	return AuthConfig{
		GlobalCfg: globalConf,
		JWTSecret: os.Getenv("JWT_SECRET"),
		JWTExpMin: os.Getenv("JWT_EXPIRATION_MINUTES"),
	}, err

}
