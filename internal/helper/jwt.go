package helpers

import (
	"aplikasi_restoran/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTClaim struct {
	ID    uint   `json:"user_id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.RegisteredClaims
}

func GenerateToken(id uint, name string, email string, role string) (string, error) {
	claims := &JWTClaim{
		ID:    id,
		Name:  name,
		Email: email,
		Role:  role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)), // Token 24 jam
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(config.GetJWTSecret())
}
