package helpers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ParseID(ctx *gin.Context, param string) (uint, bool) {
	
	missingIdMessage := "missing " + param
	idParam := ctx.Param(param)
	if idParam == "" {
		ResponseError(ctx, http.StatusBadRequest, errors.New(missingIdMessage))
		return 0, false
	}
	invalidIdMessage := "invalid " + param
	idInt, err := strconv.Atoi(idParam)
	if err != nil || idInt <= 0 {
		ResponseError(ctx, http.StatusBadRequest, errors.New(invalidIdMessage))
		return 0, false
	}

	return uint(idInt), true
}
