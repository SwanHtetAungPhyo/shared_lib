package main

import (
	"github.com/ProjectSMAA/commons/config"
	"github.com/ProjectSMAA/commons/logging"
	"github.com/ProjectSMAA/email-service/handler"
	"go.uber.org/zap"
	"os"
	"strconv"
)

func main() {
	logging.LogInit()
	port := os.Getenv("PORT")
	handler.ServerStart(port)

	loadConfig, err := config.LoadConfig("./config.toml")
	if err != nil {
		return
	}

	logging.AppLogger.Info("Code", zap.String("code", config.CodeGen()))
	logging.AppLogger.Info("Load config success")
	logging.AppLogger.Info("Email service starting...", zap.Any("config", loadConfig))
	logging.AppLogger.Info("Port", zap.String("port", strconv.Itoa(loadConfig.SMTP.Port)))
}
