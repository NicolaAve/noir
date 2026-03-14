package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// creiamo un pass valido per 72 ore
func GenerateToken(userID uint) (string, error) {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "fallback_secret_in_caso_di_emergenza"
	}

	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}
