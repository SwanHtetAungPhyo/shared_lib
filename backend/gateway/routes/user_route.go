package routes

import (
	handlers "github.com/ProjectSMAA/gateway/handler"
	"github.com/gofiber/fiber/v3"
)

func SetupUserRoutes(app *fiber.App) {
	useHandler := handlers.HandleConfig.UserHandle
	app.Get("/user", useHandler.GetUser)
	app.Post("/user", useHandler.RegisterUser)
}
