package handler

import (
	"github.com/ProjectSMAA/commons/logging"
	"github.com/ProjectSMAA/commons/protos"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
)

func ServerStart(port string) {
	list, err := net.Listen("tcp", ":"+port)
	if err != nil {
		logging.AppLogger.Fatal("failed to listen: %v", zap.Error(err))
		return
	}
	server := grpc.NewServer()
	protos.RegisterEmailServicesServer(server, &EmailServer{})
	logging.AppLogger.Info("Starting server", zap.String("port", port))
	if err := server.Serve(list); err != nil {
		logging.AppLogger.Fatal("failed to serve: %v", zap.Error(err))
	}
}
