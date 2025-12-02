package menucontroller

import (
	"aplikasi_restoran/internal/dto"
	helpers "aplikasi_restoran/internal/helper"
	"aplikasi_restoran/internal/models"
	menuservice "aplikasi_restoran/internal/services/menu"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MenuController struct {
	service menuservice.MenuService
}

func NewController(service menuservice.MenuService) *MenuController {
	return &MenuController{service}
}

func (c MenuController) AddMenu(ctx *gin.Context) {
	var input dto.AddMenuRequest
	if !helpers.BindAndValidate(ctx, &input) {
		return
	}

	menu, err := c.service.AddMenu(input.Name, models.MenuCategory(input.Category),
		input.Price, input.Description, input.ImageURL, input.IsAvailable)

	if err != nil {
		helpers.ResponseError(ctx, http.StatusBadRequest, err)
		return
	}

	helpers.ResponseSuccess(ctx, http.StatusCreated, "Success Add Menu", menu)
}

func (c MenuController) GetAllMenu(ctx *gin.Context) {
	menus, err := c.service.GetAllMenu()
	if !helpers.CheckError(ctx, err) {
		return
	}
	helpers.ResponseSuccess(ctx, http.StatusOK, "Success get all menus", menus)
}

func (c MenuController) GetMenu(ctx *gin.Context) {
	id, ok := helpers.ParseID(ctx)
	if !ok {
		return
	}

	menu, err := c.service.GetMenu(id)
	if err != nil {
		helpers.ResponseError(ctx, http.StatusNotFound, errors.New("menu not found"))
		return
	}

	helpers.ResponseSuccess(ctx, http.StatusOK, "Success get menu", menu)
}

func (c MenuController) UpdateMenu(ctx *gin.Context) {
	id, ok := helpers.ParseID(ctx)
	if !ok {
		return
	}

	var input dto.UpdateMenuRequest
	if !helpers.BindAndValidate(ctx, &input) {
		return
	}

	updated, err := c.service.UpdateMenu(id, models.MenuCategory(input.Category),
		input.Name, input.Price, input.Description, input.ImageURL, input.IsAvailable)

	if err != nil {
		helpers.ResponseError(ctx, http.StatusNotFound, errors.New("menu not found"))
		return
	}

	helpers.ResponseSuccess(ctx, http.StatusOK, "Success Update Menu", updated)
}
func (c MenuController) DeleteMenu(ctx *gin.Context) {
    id, ok := helpers.ParseID(ctx)
    if !ok {
        return
    }

    _, err := c.service.GetMenu(id)
    if err != nil {
        helpers.ResponseError(ctx, http.StatusNotFound, errors.New("menu not found"))
        return
    }

    if err := c.service.DeleteMenu(id); err != nil {
        helpers.ResponseError(ctx, http.StatusInternalServerError, err)
        return
    }

    helpers.ResponseSuccess(ctx, http.StatusOK, "Menu Deleted", nil)
}
