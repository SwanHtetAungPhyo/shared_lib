package handler

import (
	"context"
	"github.com/ProjectSMAA/commons/protos"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type EmailServer struct {
	protos.UnimplementedEmailServicesServer
}

func (s *EmailServer) SendEmail(ctx context.Context, req *protos.EmailRequest) (*protos.EmailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendEmail not implemented")
}
