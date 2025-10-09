package postgres

import (
	"log"
	"sync"

	"github.com/ahsansaif47/home-kitchens/auth/config"
	"github.com/ahsansaif47/home-kitchens/auth/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	Connection *gorm.DB
}

var dbInstance *Database
var dbOnce sync.Once

func GetDatabaseConnection() *Database {
	dbOnce.Do(func() {
		dbInstance = &Database{
			Connection: connect(),
		}
	})

	return dbInstance
}

func connect() *gorm.DB {
	c := config.GetConfig()

	db, err := gorm.Open(postgres.Open(c.DBUrl), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting database: %v", err)
	}

	err = db.AutoMigrate(
		&models.User{},
		&models.UserAddress{},
		&models.Role{},
	)

	if err != nil {
		log.Fatalf("Error migrating database: %v", err)
	}

	return db
}
