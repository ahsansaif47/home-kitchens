package main

import (
	"fmt"
	"log"

	"github.com/ahsansaif47/home-kitchens/auth/config"
	"github.com/ahsansaif47/home-kitchens/auth/gRPC/services"
	"github.com/ahsansaif47/home-kitchens/auth/http/routes"
	"github.com/ahsansaif47/home-kitchens/auth/repository/postgres"
	c "github.com/ahsansaif47/home-kitchens/common/database"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func init() {

	// TODO: Wrap in retry func
	emailRpcClient, err := services.NewEmailClient(":50052")
	if err != nil {
		log.Fatalf("failed to connect to email client: %s", err.Error())
	}

	services.EmailServiceClient = emailRpcClient

	// ...

}

func main() {

	db := postgres.GetDatabaseConnection().Connection
	cache := c.NewCache()

	// go func() {
	// 	InitializegRPCClient()
	// }()

	startHTTP(db, cache)
}

func startHTTP(db *gorm.DB, cache c.Cache) {
	app := fiber.New()
	routes.InitRoutes(app, db, &cache)

	port := config.GetConfig().GlobalCfg.Port
	log.Printf("Fiber server listening on port: %s", port)
	log.Fatal(app.Listen(fmt.Sprintf(":%s", port)))

}
