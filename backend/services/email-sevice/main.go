package main

import (
	"github.com/ProjectSMAA/commons/logging"
	"github.com/ProjectSMAA/email-service/config"
	"go.uber.org/zap"
	"strconv"
)

func main() {
	logging.LogInit()
	loadConfig, err := config.LoadConfig("./config.toml")
	if err != nil {
		return
	}

	logging.AppLogger.Info("Load config success")
	logging.AppLogger.Info("Email service starting...", zap.Any("config", loadConfig))
	logging.AppLogger.Info("Port", zap.String("port", strconv.Itoa(loadConfig.SMTP.Port)))
}
