package middlewares

import (
	helpers "aplikasi_restoran/internal/helper"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Role(role string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		roleNow := "hanya role " + role
		userRole, ok := ctx.Get("user_role")
		if !ok || userRole != role {
			helpers.ResponseError(ctx, http.StatusForbidden, errors.New(roleNow))
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
