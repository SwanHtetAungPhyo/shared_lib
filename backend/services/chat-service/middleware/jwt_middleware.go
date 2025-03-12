package middleware

import (
	"fmt"
	"github.com/ProjectSMAA/commons/logging"
	"github.com/ProjectSMAA/commons/models"
	"github.com/golang-jwt/jwt/v4"
	"go.uber.org/zap"
	"time"
)

type Claims struct {
	UserId string `json:"user_id"`
	jwt.RegisteredClaims
}

func GenerateJWT(userID string) string {
	claims := Claims{
		UserId: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "test",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(models.JwtSecret)
	if err != nil {
		logging.AppLogger.Error("Error generating JWT token", zap.Error(err))
		return ""
	}
	return tokenString
}

func WebSockMiddleWare(token string) bool {
	tokenString, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return models.JwtSecret, nil
	})
	if err != nil {
		logging.AppLogger.Error("Error parsing token", zap.Error(err))
		return false
	}
	claims, ok := tokenString.Claims.(*Claims)
	if !ok || !tokenString.Valid {
		logging.AppLogger.Error("Error parsing token", zap.Error(err))
		logging.AppLogger.Warn("Invalid token")
		return false
	}
	logging.AppLogger.Info("Token parsed")
	logging.AppLogger.Info("Authenticated User", zap.String("user_id", claims.UserId))
	return true
}
