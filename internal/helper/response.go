package helpers

import "github.com/gin-gonic/gin"

type Response struct {
	Code    int         `json:"code"`
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func emptyObject(data interface{}) interface{} {
	if data == nil {
		return gin.H{}
	}
	return data
}

func ResponseSuccess(ctx *gin.Context, code int, message string, data interface{}) {
	resp := Response{
		Code:    code,
		Status:  "success",
		Message: message,
		Data:    emptyObject(data),
	}

	ctx.JSON(code, resp)
}

func ResponseError(ctx *gin.Context, code int, err error) {
	resp := Response{
		Code:    code,
		Status:  "error",
		Message: err.Error(),
		Data:    gin.H{},
	}

	ctx.JSON(code, resp)
}
