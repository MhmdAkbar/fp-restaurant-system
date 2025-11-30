package middlewares

import (
	"aplikasi_restoran/config"
	helpers "aplikasi_restoran/internal/helper"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		auth := ctx.GetHeader("Authorization")
		if auth == "" || !strings.HasPrefix(auth, "Bearer ") {
			helpers.ResponseError(ctx, http.StatusUnauthorized, errors.New("token required"))
			ctx.Abort()
			return
		}

		tokenString := strings.TrimPrefix(auth, "Bearer ")

		claims := &helpers.JWTClaim{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
			return config.GetJWTSecret(), nil
		})

		if err != nil || !token.Valid {
			helpers.ResponseError(ctx, http.StatusUnauthorized, errors.New("invalid token"))
			ctx.Abort()
			return
		}

		// Simpan data claim ke context (TANPA float64 lagi)
		ctx.Set("user_id", claims.Id)
		ctx.Set("user_role", claims.Role)
		fmt.Println("User ID from token auth:", claims.Id)

		ctx.Next()
	}
}
