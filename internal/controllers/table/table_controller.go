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

// type UserController struct {
// 	service userservice.UserService
// }

// func NewController(service userservice.UserService) *UserController {
// 	return &UserController{service}
// }

type TableController struct {
	service tableservice.TableService
}

func NewController(service tableservice.TableService) *TableController {
	return &TableController{service}
}

func (c TableController) AddTable(ctx *gin.Context) {
	var input dto.AddTableRequest

	if err := ctx.ShouldBindJSON(&input); err != nil {
		helpers.ResponseError(ctx, http.StatusBadRequest, err)
		return
	}

	table, err := c.service.AddTable(input.TableNumber, input.QrCode, input.Available)
	if err != nil {
		helpers.ResponseError(ctx, http.StatusBadRequest, err)
		return
	}

	result := dto.AddTableResponse{
		TableNumber: table.TableNumber,
		QrCode:      table.QrCode,
		Available:   table.Status,
	}

	helpers.ResponseSuccess(ctx, http.StatusOK, "Success Add Table", result)

}

func (c TableController) GetTable(ctx *gin.Context) {
	idParam := ctx.Param("id")

	if idParam == "" {
		helpers.ResponseError(ctx, http.StatusBadRequest, errors.New("param is empty"))
		return
	}
	var id uint
	_, err := fmt.Sscan(idParam, &id)
	if err != nil {
		helpers.ResponseError(ctx, http.StatusBadRequest, err)
		return
	}

	table, err := c.service.GetTable(id)
	if err != nil {
		helpers.ResponseError(ctx, http.StatusNotFound, err)
		return
	}

	helpers.ResponseSuccess(ctx, http.StatusOK, "Success get table", table)

}

func (c TableController) UpdateTable (ctx *gin.Context) {
	var input dto.AddTableRequest

	if err := ctx.ShouldBindJSON(&input); err != nil {
		helpers.ResponseError(ctx, http.StatusBadRequest, err)
		return
	}

	table, err := c.service.UpdateTable(input.Id,input.TableNumber, input.QrCode, input.Available)
	fmt.Println(input.Id)
	fmt.Println(input.TableNumber)
	fmt.Println(input.QrCode)
	fmt.Println(input.Available)
	if err != nil {
		helpers.ResponseError(ctx, http.StatusBadRequest, err)
		return
	}

	result := dto.AddTableResponse{
		TableNumber: table.TableNumber,
		QrCode:      table.QrCode,
		Available:   table.Status,
	}
	fmt.Println(result)

	helpers.ResponseSuccess(ctx, http.StatusOK, "Success update Table", result)
}

func (c TableController)UpdateStatus (ctx *gin.Context) {
	var input dto.AddTableRequest

	if err := ctx.ShouldBindJSON(&input); err != nil {
		helpers.ResponseError(ctx, http.StatusBadRequest, err)
		return
	}

	table, err := c.service.UpdateStatus(input.Id, input.Available )
	if err != nil {
		helpers.ResponseError(ctx, http.StatusBadRequest, err)
		return
	}
	fmt.Println("avaliable dari controller input :", input.Available)
	fmt.Println("status dari controller table :", table.Status)

	result := dto.AddTableResponse{
		TableNumber: table.TableNumber,
		QrCode:      table.QrCode,
		Available:   table.Status,
	}

	helpers.ResponseSuccess(ctx, http.StatusOK, "Success update Table", result)

}

func (c TableController) DeleteTable (ctx *gin.Context) {
	var input dto.AddTableRequest

	if err := ctx.ShouldBindJSON(&input); err != nil {
		helpers.ResponseError(ctx,http.StatusBadRequest, err)
	}

	err := c.service.DeleteTable(input.Id)
	if err != nil  {
		helpers.ResponseError(ctx, http.StatusBadRequest,err)
	} 

	helpers.ResponseSuccess(ctx, http.StatusOK, "success delete table", nil)
}
