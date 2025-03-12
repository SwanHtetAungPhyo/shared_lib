package config

import (
	"bytes"
	"github.com/ProjectSMAA/commons/logging"
	"github.com/ProjectSMAA/commons/models"
	"go.uber.org/zap"
	"html/template"
)

func LoadTemplate(filepath string, email *models.Email) (string, error) {
	tmpl, err := template.ParseFiles(filepath)
	if err != nil {
		logging.AppLogger.Error("Error loading template", zap.Error(err))
		return "", err
	}
	var buffer bytes.Buffer
	err = tmpl.Execute(&buffer, email)
	if err != nil {
		logging.AppLogger.Error("Error loading template", zap.Error(err))
		return "", err
	}
	return buffer.String(), nil
}
