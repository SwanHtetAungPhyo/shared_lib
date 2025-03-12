package routes

import (
	"github.com/ProjectSMAA/gateway/http/handler"
	"github.com/gofiber/fiber/v3"
)

func SettingAppRoute(app *fiber.App) {
	SetupUserRoutes(app)
	SetUpServiceRoutes(app)
}
func SetupUserRoutes(app *fiber.App) {
	useHandler := handlers.HandleConfig.UserHandle
	userRoutes := app.Group("/user")
	userRoutes.Get("/user", useHandler.GetUser)
	userRoutes.Post("/user", useHandler.RegisterUser)
}

func SetUpServiceRoutes(app *fiber.App) {
	serviceHandle := handlers.HandleConfig.ServiceBlockHandle
	serviceRoute := app.Group("/service")
	serviceRoute.Post("/", serviceHandle.CreateServiceOnPlatform)
	serviceRoute.Get("/", serviceHandle.ListServiceBlocks)
	serviceRoute.Put("/", serviceHandle.UpdateServiceOnPlatform)
	serviceRoute.Delete("/", serviceHandle.DeleteServiceOnPlatform)
}
