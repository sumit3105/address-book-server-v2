package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(jwtSecret string, userID uint64, userEmail string) (string, error) {
	
	claims := jwt.MapClaims{
		"user_id": userID,
		"user_email": userEmail,
		"exp": time.Now().Add(24 * time.Hour).Unix(),
	}

	token:=jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jwtSecret))
}