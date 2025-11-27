package middlewares

import (
	helpers "aplikasi_restoran/internal/helper"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func VerifyUserAccess() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		// Ambil user_id dari context
		uidValue, exists := ctx.Get("user_id")
		if !exists {
			helpers.ResponseError(ctx, http.StatusUnauthorized, errors.New("user id not found in token"))
			ctx.Abort()
			return
		}

		uidFromToken, ok := uidValue.(uint)
		if !ok {
			helpers.ResponseError(ctx, http.StatusUnauthorized, errors.New("invalid user id type"))
			ctx.Abort()
			return
		}

		// Ambil role dari context
		roleValue, exists := ctx.Get("user_role")
		if !exists {
			helpers.ResponseError(ctx, http.StatusUnauthorized, errors.New("role not found in token"))
			ctx.Abort()
			return
		}

		roleFromToken, ok := roleValue.(string)
		if !ok {
			helpers.ResponseError(ctx, http.StatusUnauthorized, errors.New("invalid role type"))
			ctx.Abort()
			return
		}

		// Jika super_admin, izinkan akses seluruh user
		if roleFromToken == "super_admin" {
			ctx.Next()
			return
		}

		// Ambil id dari param
		paramID, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			helpers.ResponseError(ctx, http.StatusBadRequest, errors.New("invalid user id"))
			ctx.Abort()
			return
		}

		// Bandingkan ID token vs ID request param
		if uint(paramID) != uidFromToken {
			helpers.ResponseError(ctx, http.StatusForbidden,
				errors.New("access forbidden: cannot access another user's profile"))
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
