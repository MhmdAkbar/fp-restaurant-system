package helpers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// BindAndValidate digunakan untuk binding dan validasi input JSON
func BindAndValidate(ctx *gin.Context, input interface{}) bool {
	if err := ctx.ShouldBindJSON(input); err != nil {
		ResponseError(ctx, http.StatusBadRequest, err)
		return false
	}
	return true
}

// CheckError digunakan untuk men-handle error umum di controller
func CheckError(ctx *gin.Context, err error) bool {
	if err != nil {
		ResponseError(ctx, http.StatusInternalServerError, err)
		return false
	}
	return true
}
