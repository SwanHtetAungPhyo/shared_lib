package main

import (
	"github.com/ProjectSMAA/auth-service/internal/handler"
	"github.com/ProjectSMAA/auth-service/internal/repo"
	"github.com/ProjectSMAA/auth-service/internal/services"
	"github.com/ProjectSMAA/commons/database"
	"github.com/ProjectSMAA/commons/logging"
	"log"
	"net"

	"github.com/ProjectSMAA/commons/protos" // Import your generated protobuf package
	"google.golang.org/grpc"
)

func main() {
	logging.LogInit()
	defer logging.ShutdownLogger()
	database.ConnectDB()
	database.Migration()
	repo := repo.NewUserRepo()
	service := services.NewAuthService(repo)
	authServer := handler.NewServer(service)
	listener, err := net.Listen("tcp", ":5001")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	server := grpc.NewServer()
	protos.RegisterUserServiceServer(server, authServer)

	log.Println("gRPC server is running on port :5001...")
	if err := server.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
