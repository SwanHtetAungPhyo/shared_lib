package routes

import (
	handlers "github.com/ProjectSMAA/gateway/handler"
	"github.com/gofiber/fiber/v3"
)

// SetupUserRoutes defines user-related routes
func SetupUserRoutes(app *fiber.App) {
	app.Get("/user", handlers.GetUser)
	app.Post("/user", handlers.RegisterUser)
}
