package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Secret key for signing JWT tokens
var jwtSecret = []byte("DyxYtuA5nIN9kCWlGaUgYC8sm2215fgzxYj+XA8l/T4=")

// Replace with environment variable in production

// GenerateJWT generates a JWT token for the given user ID
func GenerateJWT(userID string) (string, error) {

	expirationTime := time.Now().Add(24 * time.Hour)

	claims := &jwt.MapClaims{
		"user_id": userID,
		"exp":     expirationTime.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(jwtSecret)
}
