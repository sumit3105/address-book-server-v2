package utils

import (
	"address-book-server-v2/internal/core/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(serverCfg *config.ServerConfig, userID uint, userEmail string) (string, error) {
	
	jwtSecret := serverCfg.JwtSecret

	claims := jwt.MapClaims{
		"user_id": userID,
		"user_email": userEmail,
		"exp": time.Now().Add(24 * time.Hour).Unix(),
	}

	token:=jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jwtSecret))
}