package tablecontroller

import (
	"aplikasi_restoran/internal/dto"
	helpers "aplikasi_restoran/internal/helper"
	tableservice "aplikasi_restoran/internal/services/table"
	"net/http"
	"strconv"
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

	if !helpers.BindAndValidate(ctx, &input) {
		return
	}

	table, err := c.service.AddTable(input.QrCode, input.Status)
	if !helpers.CheckError(ctx, err) {
		return
	}

	response := dto.TableResponse{
		Id:        table.ID,
		TableCode: table.TableCode,
		QrCode:    table.QrCode,
		Status:    table.Status,
	}

	helpers.ResponseSuccess(ctx, http.StatusCreated, "Success Add Table", response)
}

func (c TableController) GetTable(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if !helpers.CheckError(ctx, err) {
		return
	}

	table, err := c.service.GetTable(uint(id))
	if !helpers.CheckError(ctx, err) {
		return
	}

	response := dto.TableResponse{
		Id:        table.ID,
		TableCode: table.TableCode,
		QrCode:    table.QrCode,
		Status:    table.Status,
	}

	helpers.ResponseSuccess(ctx, http.StatusOK, "Success Get Table", response)
}

func (c TableController) GetAll(ctx *gin.Context) {
	tables, err := c.service.GetAll()
	if !helpers.CheckError(ctx, err) {
		return
	}

	var result []dto.TableResponse
	for _, t := range tables {
		result = append(result, dto.TableResponse{
			Id:        t.ID,
			TableCode: t.TableCode,
			QrCode:    t.QrCode,
			Status:    t.Status,
		})
	}

	helpers.ResponseSuccess(ctx, http.StatusOK, "Success Get All Tables", result)
}

func (c TableController) UpdateTable(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if !helpers.CheckError(ctx, err) {
		return
	}

	var input dto.TableRequest
	if !helpers.BindAndValidate(ctx, &input) {
		return
	}

	table, err := c.service.UpdateTable(uint(id), input.QrCode, input.Status)
	if !helpers.CheckError(ctx, err) {
		return
	}

	response := dto.TableResponse{
		Id:        table.ID,
		TableCode: table.TableCode,
		QrCode:    table.QrCode,
		Status:    table.Status,
	}

	helpers.ResponseSuccess(ctx, http.StatusOK, "Success Update Table", response)
}

func (c TableController) UpdateStatus(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if !helpers.CheckError(ctx, err) {
		return
	}

	var body struct {
		Status bool `json:"status" binding:"required"`
	}

	if !helpers.BindAndValidate(ctx, &body) {
		return
	}

	table, err := c.service.UpdateStatus(uint(id), body.Status)
	if !helpers.CheckError(ctx, err) {
		return
	}

	response := dto.TableResponse{
		Id:        table.ID,
		TableCode: table.TableCode,
		QrCode:    table.QrCode,
		Status:    table.Status,
	}

	helpers.ResponseSuccess(ctx, http.StatusOK, "Status Updated", response)
}
