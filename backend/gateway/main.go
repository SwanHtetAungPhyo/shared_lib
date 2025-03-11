package main

import (
	"context"
	"log"

	"github.com/gofiber/fiber/v3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/ProjectSMAA/commons/models"
	"github.com/ProjectSMAA/commons/protos" // Import the generated gRPC package
)

// gRPC Client Function to avoid duplication
func grpcClient() (*grpc.ClientConn, error) {
	return grpc.Dial("localhost:5001", grpc.WithTransportCredentials(insecure.NewCredentials()))
}

// gRPC Call Helper Function
func callGrpcService[T any](clientFunc func(protos.UserServiceClient) (T, error)) (T, error) {
	var result T
	conn, err := grpcClient()
	if err != nil {
		return result, err
	}
	defer conn.Close()

	client := protos.NewUserServiceClient(conn)
	return clientFunc(client)
}

// RegisterUser handles user registration
func RegisterUser(c fiber.Ctx) error {
	var user *models.User
	if err := c.Bind().Body(&user); err != nil {
		return c.Status(400).SendString("Invalid request")
	}

	// Call gRPC service to register user
	resp, err := callGrpcService(func(client protos.UserServiceClient) (*protos.UserRegisteredResponse, error) {
		return client.RegisterUser(context.Background(), &protos.UserRegisterRequest{
			Email:    user.Email,
			Username: user.Username,
			Password: "<PASSWORD>", // Ensure to secure password handling
		})
	})
	if err != nil {
		return c.Status(500).SendString("gRPC error: " + err.Error())
	}

	return c.Status(200).JSON(fiber.Map{"id": resp.Id, "username": resp.Username, "email": resp.Email, "status": resp.Status})
}

// GetUser checks if a user exists based on the email
func getUser(c fiber.Ctx) error {
	// Call gRPC service to check user existence
	resp, err := callGrpcService(func(client protos.UserServiceClient) (*protos.UserResponse, error) {
		return client.CheckUserExistance(context.Background(), &protos.UserRequest{Email: "John Doe"})
	})
	if err != nil {
		return c.Status(500).SendString("gRPC error: " + err.Error())
	}

	// Return response
	return c.JSON(fiber.Map{"user_exists": resp.Exists})
}

func main() {
	app := fiber.New()

	// Define API routes
	app.Get("/user", getUser)
	app.Post("/user", RegisterUser)

	// Start Fiber server
	if err := app.Listen(":3000"); err != nil {
		log.Fatal(err)
	}
}
