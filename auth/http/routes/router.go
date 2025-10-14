package routes

import (
	// "github.com/ahsansaif47/home-kitchens/auth/config"
	"fmt"

	"github.com/ahsansaif47/home-kitchens/auth/http/controllers"
	"github.com/ahsansaif47/home-kitchens/auth/repository/postgres"
	r "github.com/ahsansaif47/home-kitchens/auth/repository/redis"
	"github.com/gofiber/fiber/v2"

	"github.com/gofiber/swagger"
	"gorm.io/gorm"
)

func InitRoutes(app *fiber.App) {
	// config := config.GetConfig()

	app.Get("/swagger/*", swagger.HandlerDefault)
	db := postgres.GetDatabaseConnection().Connection
	cache := r.NewCache()

	api := app.Group("/api")

	v1 := api.Group("/v1")

	userRoutes := v1.Group("/users")
	InitUserRoutes(userRoutes, db, cache)
}

func InitUserRoutes(userRoutes fiber.Router, db *gorm.DB, cache r.Cache) {
	// TODO: Implement the repository then the serice and then handler

	userRepo := postgres.NewUserRepository(db)
	userService := controllers.NewUserService(userRepo)

	fmt.Println(userService)
	// userHandler := handlers.NewUserHandler(userService)

	userRoutes.Get("/", func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusAccepted)
	})
}
