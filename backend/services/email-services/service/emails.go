package service

import (
	"github.com/ProjectSMAA/commons/config"
	"github.com/ProjectSMAA/commons/logging"
	"go.uber.org/zap"
	"gopkg.in/gomail.v2"
)

type EmailService struct {
	config *config.Config
}

func NewEmailService(cfg *config.Config) *EmailService {
	return &EmailService{
		config: cfg,
	}
}

func (s EmailService) SendCode(code, to string, subject string) error {
	mail := gomail.NewMessage()
	mail.SetHeader("From", s.config.SMTP.Email)
	mail.SetHeader("To", to)
	mail.SetHeader("Subject", subject)
	mail.SetBody("text/html", code)
	dialing := gomail.NewDialer(s.config.SMTP.Protocol, s.config.SMTP.Port, s.config.SMTP.Email, s.config.SMTP.Password)
	if err := dialing.DialAndSend(mail); err != nil {
		logging.AppLogger.Error("Send Email Failed", zap.Error(err))
		return err
	}
	logging.AppLogger.Info("Send Email Success", zap.String("to", to))
	return nil
}
