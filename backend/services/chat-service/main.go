package main

import (
	"github.com/ProjectSMAA/commons/logging"
	chathandler "github.com/SwanHtetAungPhyo/chat/handler"
	"go.uber.org/zap"
	"net/http"
)

func main() {
	logging.LogInit()
	http.HandleFunc("/ws", chathandler.HandleConnection)
	http.HandleFunc("/token", chathandler.TokenHandler)
	logging.AppLogger.Info("Trying to start Server")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		logging.AppLogger.Error("Error starting chat server", zap.Error(err))
	}
	logging.AppLogger.Info("Server started")
}
