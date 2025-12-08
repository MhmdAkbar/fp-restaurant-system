package helpers

import (
	"aplikasi_restoran/config"
	"time"
	"github.com/golang-jwt/jwt/v5"
)

type JWTClaim struct {
	Id uint `json:"user_id"`
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.RegisteredClaims
}

func GenerateToken(id uint, name string, email string, role string) (string, error) {
	claims := &JWTClaim{
		Id: id,
		Email: email,
		Role:  role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 240)), // Token 24 jam // pindahkan ke dot env
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(config.GetJWTSecret())
}
