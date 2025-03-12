package repo

import (
	"github.com/ProjectSMAA/commons/database"
	"github.com/ProjectSMAA/commons/logging"
	"github.com/ProjectSMAA/commons/models"
	"go.uber.org/zap"
)

type UserRepo struct {
}

func NewUserRepo() *UserRepo {
	return &UserRepo{}
}

func (r *UserRepo) Register(req *models.User) {
	logging.AppLogger.Info("Registering user")
	if err := database.DB.Create(req).Error; err != nil {
		logging.AppLogger.Error("Error in  creating user", zap.String("error", err.Error()))
		panic(err)
	}
}
