package converter

import (
	"github.com/ProjectSMAA/commons/models"
	"github.com/ProjectSMAA/commons/protos"
)

// http models.user -> proto
// proto -> http
type UserConverter struct{}

func NewUserConverter() *UserConverter {
	return &UserConverter{}
}

func (uc *UserConverter) HttpToProto(user *models.User) *protos.UserRegisterRequest {
	return &protos.UserRegisterRequest{
		Email:    user.Email,
		Username: user.Username,
		Password: "<PASSWORD>",
	}
}

func (uc *UserConverter) ProtoToHttp(user *protos.UserRegisterRequest) *models.User {
	return &models.User{
		Email:    user.Email,
		Username: user.Username,
		Status:   "active",
	}
}
