package handlers

import (
	"context"
	"github.com/ProjectSMAA/commons/models"
	"github.com/ProjectSMAA/commons/protos"
	"github.com/ProjectSMAA/gateway/grpc"
	"github.com/ProjectSMAA/gateway/utils"

	"github.com/gofiber/fiber/v3"
)

type UserHandlerInterface interface {
	RegisterUser(c fiber.Ctx) error
	GetUser(c fiber.Ctx) error
	CheckUser(c fiber.Ctx) error
}
type UsersHandler struct {
	UserHandler UserHandlerInterface
}

func NewUsersHandler() *UsersHandler {
	return &UsersHandler{}
}

func (u *UsersHandler) RegisterUser(c fiber.Ctx) error {
	var user models.User
	if err := c.Bind().Body(&user); err != nil {
		return c.Status(400).SendString("Invalid request")
	}

	resp, err := grpc.CallGrpcService("localhost",
		"50051",
		protos.NewUserServiceClient,
		func(client protos.UserServiceClient) (*protos.UserRegisteredResponse, error) {
			return client.RegisterUser(context.Background(), &protos.UserRegisterRequest{
				Email:    user.Email,
				Username: user.Username,
				Password: "<PASSWORD>",
			})
		})
	token, err := utils.GenerateJWT(resp.Email)
	if err != nil {
		return c.Status(500).SendString("Token generation failed")
	}
	if err != nil {
		return c.Status(500).SendString("gRPC error: " + err.Error())
	}
	refresh := fiber.Cookie{
		Name:  "refresh_token",
		Value: token,
	}
	c.Cookie(&refresh)
	return c.Status(200).JSON(fiber.Map{"id": resp.Id,
		"username": resp.Username,
		"email":    resp.Email,
		"status":   resp.Status, "token": token})
}

func (us *UsersHandler) GetUser(c fiber.Ctx) error {
	resp, err := grpc.CallGrpcService(
		"localhost",
		"50051",
		protos.NewUserServiceClient,
		func(client protos.UserServiceClient) (*protos.UserResponse, error) {
			return client.CheckUserExistance(context.Background(), &protos.UserRequest{
				Email: c.Params("email"),
			})
		},
	)
	if err != nil {
		return c.Status(500).SendString("gRPC error: " + err.Error())
	}

	return c.JSON(fiber.Map{"user_exists": resp.Exists})
}

func (us *UsersHandler) CheckUser(c fiber.Ctx) error {
	return c.Status(200).SendString("OK")
}
