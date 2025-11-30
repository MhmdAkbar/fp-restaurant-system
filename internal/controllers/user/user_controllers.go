package usercontroller

import (
	"aplikasi_restoran/internal/dto"
	helpers "aplikasi_restoran/internal/helper"
	usermodels "aplikasi_restoran/internal/models/user"
	userservice "aplikasi_restoran/internal/services/user"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	service userservice.UserService
}

func NewController(service userservice.UserService) *UserController {
	return &UserController{service}
}

func (c *UserController) Register(ctx *gin.Context) {
	var input dto.RegisterRequest

	if err := ctx.ShouldBindJSON(&input); err != nil {
		helpers.ResponseError(ctx, http.StatusBadRequest, err)
		return
	}

	// Dapatkan user setelah insert
	user, err := c.service.Register(input.Name, input.Email, input.Password, usermodels.UserRole(input.Role))
	if err != nil {
		helpers.ResponseError(ctx, http.StatusBadRequest, err)
		return
	}

	result := dto.RegisterResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Role:  string(user.Role),
	}

	helpers.ResponseSuccess(ctx, http.StatusOK, "register sukses", result)
}

func (c *UserController) Login(ctx *gin.Context) {
	var input dto.LoginRequest
	if err := ctx.ShouldBindJSON(&input); err != nil {
		helpers.ResponseError(ctx, http.StatusBadRequest, err)
		return
	}

	user, err := c.service.Login(input.Email, input.Password)
	if err != nil {
		helpers.ResponseError(ctx, http.StatusUnauthorized, err)
		return
	}
	token, err := helpers.GenerateToken(user.ID, user.Name, user.Email, string(user.Role))
	if err != nil {
		helpers.ResponseError(ctx, http.StatusInternalServerError, err)
	}
	result := dto.LoginResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Role:  string(user.Role),
		Token: token,
	}

	helpers.ResponseSuccess(ctx, http.StatusOK, "success login", result)
}

func (c *UserController) GetProfile(ctx *gin.Context) {
	idParam := ctx.Param("id")

	if idParam == "" {
		helpers.ResponseError(ctx, http.StatusBadRequest, fmt.Errorf("id wajib diisi"))
		return
	}

	var id uint
	_, err := fmt.Sscan(idParam, &id)
	if err != nil {
		helpers.ResponseError(ctx, http.StatusBadRequest, fmt.Errorf("id tidak valid"))
		return
	}

	user, err := c.service.GetProfile(id)
	if err != nil {
		helpers.ResponseError(ctx, http.StatusNotFound, err)
		return
	}

	helpers.ResponseSuccess(ctx, http.StatusOK, "success", user)
}

func (c *UserController) UpdateProfile(ctx *gin.Context) {
	idParam := ctx.Param("id")

	if idParam == "" {
		helpers.ResponseError(ctx, http.StatusBadRequest, fmt.Errorf("id wajib diisi"))
		return
	}

	var id uint
	_, err := fmt.Sscan(idParam, &id)
	if err != nil {
		helpers.ResponseError(ctx, http.StatusBadRequest, fmt.Errorf("id tidak valid"))
		return
	}

	var input dto.UpdateProfileRequest
	if err := ctx.ShouldBindJSON(&input); err != nil {
		helpers.ResponseError(ctx, http.StatusBadRequest, err)
		return
	}

	updated, err := c.service.UpdateProfile(id, input.Name, input.Email)
	if err != nil {
		helpers.ResponseError(ctx, http.StatusInternalServerError, err)
		return
	}

	result := dto.UpdateProfileResponse{
		ID:    updated.ID,
		Name:  updated.Name,
		Email: updated.Email,
	}

	helpers.ResponseSuccess(ctx, http.StatusOK, "profile updated successfully", result)
}

func (c *UserController) DeleteProfile(ctx *gin.Context) {
	idParam := ctx.Param("id")

	if idParam == "" {
		helpers.ResponseError(ctx, http.StatusBadRequest, fmt.Errorf("id wajib diisi"))
		return
	}

	var id uint
	_, err := fmt.Sscan(idParam, &id)
	if err != nil {
		helpers.ResponseError(ctx, http.StatusBadRequest, fmt.Errorf("id tidak valid"))
		return
	}
	c.service.DeleteProfile(id)
	helpers.ResponseSuccess(ctx, http.StatusOK, "success", nil)
}
