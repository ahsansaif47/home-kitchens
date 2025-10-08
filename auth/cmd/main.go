package cmd

import (
	"fmt"

	"github.com/ahsansaif47/home-kitchens/auth/config"
	"github.com/ahsansaif47/home-kitchens/auth/http/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()

	routes.InitRoutes(app)

	app.Listen(fmt.Sprintf(":", config.GetConfig().Port))

}
