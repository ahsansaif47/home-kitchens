package routes

import (
	// "github.com/ahsansaif47/home-kitchens/auth/config"

	"github.com/ahsansaif47/home-kitchens/auth/gRPC/services"
	"github.com/ahsansaif47/home-kitchens/auth/http/controllers"
	"github.com/ahsansaif47/home-kitchens/auth/http/handlers"
	"github.com/ahsansaif47/home-kitchens/auth/repository/postgres"
	"github.com/ahsansaif47/home-kitchens/auth/repository/redis"
	"github.com/gofiber/fiber/v2"

	_ "github.com/ahsansaif47/home-kitchens/auth/docs" // 👈 this is important
	"github.com/gofiber/swagger"
	"gorm.io/gorm"
)

// @title						HomeKitchens Local API
// @version					1.0
// @description				This is a swagger for HomeKitchens
// @host						localhost:8080
// @BasePath					/api/v1
// @schemes					http
// @securityDefinitions.apikey	BearerAuth
// @in							header
// @name						Authorization
func InitRoutes(app *fiber.App, db *gorm.DB, cache redis.ICacheRepository) {
	app.Get("/swagger/*", swagger.HandlerDefault)

	api := app.Group("/api")
	v1 := api.Group("/v1")

	userRoutes := v1.Group("/users")
	InitUserRoutes(userRoutes, db, cache)
}

func InitUserRoutes(userRoutes fiber.Router, db *gorm.DB, cache redis.ICacheRepository) {
	userRepo := postgres.NewUserRepository(db)
	userService := controllers.NewUserService(userRepo, cache, *services.EmailServiceClient)
	userHandlers := handlers.NewAuthHandler(userService)

	userRoutes.Post("/signup", userHandlers.CreateUser)
	userRoutes.Post("/signin", userHandlers.Signin)
}
