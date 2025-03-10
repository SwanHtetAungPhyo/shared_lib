package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var SMMALogger *zap.SugaredLogger

func InitLogger() {
	config := zap.NewProductionConfig()
	config.Encoding = "json"
	config.Level = zap.NewAtomicLevelAt(zapcore.InfoLevel)
	config.OutputPaths = []string{
		"stdout",
	}

	logger, err := config.Build()
	if err != nil {
		panic("Failed to initialize logger: " + err.Error())
	}

	SMMALogger = logger.Sugar()

	SMMALogger.Info("Logger initialized successfully")
}

func CloseLogger() {
	SMMALogger.Sync()
}
