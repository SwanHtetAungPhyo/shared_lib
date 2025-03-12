package main

import (
	"github.com/ProjectSMAA/commons/logging"
	"github.com/ProjectSMAA/commons/util"
	config_gate "github.com/ProjectSMAA/gateway/config"
	"github.com/ProjectSMAA/gateway/http/handler"
	"github.com/ProjectSMAA/gateway/http/routes"
	"log"

	"github.com/gofiber/fiber/v3"
)

func main() {

	app := fiber.New()
	logging.LogInit()
	logging.AppLogger.Info("Logger init in the Gateway Complete")
	handlers.Init()
	logging.AppLogger.Info("Handler init in the Gateway Complete")
	routes.SettingAppRoute(app)
	logging.AppLogger.Info("Router init in the Gateway Complete")

	logging.AppLogger.Info("Gateway init in the Gateway Complete")

	logging.AppLogger.Info("Gateway is listening on the http://localhost:3000")
	app.Get("/", HealthCheck)
	if err := app.Listen(":3000"); err != nil {
		log.Fatal(err)
	}
}

func HealthCheck(c fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(util.ApiResp(fiber.StatusOK,
		"Hello from Gateway ",
		config_gate.NewGateWayRoutes(),
	))
}
