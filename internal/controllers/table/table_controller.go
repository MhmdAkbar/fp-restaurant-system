package tablecontroller

import (
	"aplikasi_restoran/internal/dto"
	helpers "aplikasi_restoran/internal/helper"
	tableservice "aplikasi_restoran/internal/services/table"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TableController struct {
	service tableservice.TableService
}

func NewController(service tableservice.TableService) *TableController {
	return &TableController{service}
}

func (c TableController) AddTable(ctx *gin.Context) {
	var input dto.TableRequest
	if err := ctx.ShouldBindJSON(&input); err != nil {
		helpers.ResponseError(ctx, http.StatusBadRequest, err)
		return
	}

	table, err := c.service.AddTable(input.QrCode, input.Status)
	if err != nil {
		helpers.ResponseError(ctx, http.StatusBadRequest, err)
		return
	}

	helpers.ResponseSuccess(ctx, http.StatusOK, "Success Add Table", table)
}

func (c TableController) GetTable(ctx *gin.Context) {
	idParam := ctx.Param("id")
	if idParam == "" {
		helpers.ResponseError(ctx, http.StatusBadRequest, errors.New("missing id"))
		return
	}

	var id uint
	fmt.Sscan(idParam, &id)

	table, err := c.service.GetTable(id)
	if err != nil {
		helpers.ResponseError(ctx, http.StatusNotFound, err)
		return
	}

	helpers.ResponseSuccess(ctx, http.StatusOK, "Success Get Table", table)
}

func (c TableController) GetAll(ctx *gin.Context) {
	tables, err := c.service.GetAll()
	if err != nil {
		helpers.ResponseError(ctx, http.StatusInternalServerError, err)
		return
	}

	helpers.ResponseSuccess(ctx, http.StatusOK, "Success Get All Tables", tables)
}

func (c TableController) UpdateTable(ctx *gin.Context) {
	idParam := ctx.Param("id")
	var id uint
	fmt.Sscan(idParam, &id)

	var input dto.TableRequest
	if err := ctx.ShouldBindJSON(&input); err != nil {
		helpers.ResponseError(ctx, http.StatusBadRequest, err)
		return
	}

	table, err := c.service.UpdateTable(id, input.QrCode, input.Status)
	if err != nil {
		helpers.ResponseError(ctx, http.StatusBadRequest, err)
		return
	}

	helpers.ResponseSuccess(ctx, http.StatusOK, "Success Update Table", table)
}

func (c TableController) UpdateStatus(ctx *gin.Context) {
	idParam := ctx.Param("id")
	var id uint
	fmt.Sscan(idParam, &id)

	var body struct {
		Status bool `json:"status"`
	}
	if err := ctx.ShouldBindJSON(&body); err != nil {
		helpers.ResponseError(ctx, http.StatusBadRequest, err)
		return
	}

	table, err := c.service.UpdateStatus(id, body.Status)
	if err != nil {
		helpers.ResponseError(ctx, http.StatusBadRequest, err)
		return
	}

	helpers.ResponseSuccess(ctx, http.StatusOK, "Status Updated", table)
}
