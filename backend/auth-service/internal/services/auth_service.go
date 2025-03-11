package services

import (
	"github.com/ProjectSMAA/auth-service/internal/repo"
	"github.com/ProjectSMAA/commons/models"
)

type AuthService struct {
	userRepo *repo.UserRepo
}

func NewAuthService(userRepo *repo.UserRepo) *AuthService {
	return &AuthService{
		userRepo: userRepo,
	}
}

func (s *AuthService) Register(req *models.User) error {
	s.userRepo.Register(req)
	return nil
}
func (s *AuthService) Login() {

}
