package helpers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ParseID(ctx *gin.Context) (uint, bool) {
	idParam := ctx.Param("id")
	if idParam == "" {
		ResponseError(ctx, http.StatusBadRequest, errors.New("missing id"))
		return 0, false
	}

	idInt, err := strconv.Atoi(idParam)
	if err != nil || idInt <= 0 {
		ResponseError(ctx, http.StatusBadRequest, errors.New("invalid id"))
		return 0, false
	}

	return uint(idInt), true
}
