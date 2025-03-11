package handler

import (
	"github.com/ProjectSMAA/auth-service/internal/services"
	"github.com/ProjectSMAA/commons/converter"
	"github.com/ProjectSMAA/commons/protos"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	protos.UnimplementedUserServiceServer
	authService *services.AuthService
}

func NewServer(authService *services.AuthService) *Server {
	return &Server{
		authService: authService,
	}
}

func (s *Server) CheckUserExistance(ctx context.Context, req *protos.UserRequest) (*protos.UserResponse, error) {
	if req.Email == "" {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid request: Name is required")
	}
	return &protos.UserResponse{
		Exists: true,
	}, nil
}

func (s *Server) RegisterUser(ctx context.Context, req *protos.UserRegisterRequest) (*protos.UserRegisteredResponse, error) {
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid request")
	}
	registeredUser := converter.NewUserConverter().ProtoToHttp(req)

	err := s.authService.Register(registeredUser)
	if err != nil {
		return nil, err
	}
	return &protos.UserRegisteredResponse{
		Id:       "1",
		Username: registeredUser.Username,
		Email:    registeredUser.Email,
		Status:   "active",
	}, nil
}
