package main

import (
	"github.com/ProjectSMAA/gateway/routes"
	"log"

	"github.com/gofiber/fiber/v3"
)

func main() {
	app := fiber.New()

	routes.SetupUserRoutes(app)

	if err := app.Listen(":3000"); err != nil {
		log.Fatal(err)
	}
}
