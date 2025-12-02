package usercontroller

import (
	"aplikasi_restoran/internal/dto"
	helpers "aplikasi_restoran/internal/helper"
	"aplikasi_restoran/internal/models"
	userservice "aplikasi_restoran/internal/services/user"
	"net/http"
	"strconv"

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

	if !helpers.BindAndValidate(ctx, &input) {
		return
	}

	user, err := c.service.Register(input.Name, input.Email, input.Password, models.UserRole(input.Role))
	if !helpers.CheckError(ctx, err) {
		return
	}

	result := dto.RegisterResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Role:  string(user.Role),
	}

	helpers.ResponseSuccess(ctx, http.StatusCreated, "register sukses", result)
}

func (c *UserController) Login(ctx *gin.Context) {
	var input dto.LoginRequest

	if !helpers.BindAndValidate(ctx, &input) {
		return
	}

	user, err := c.service.Login(input.Email, input.Password)
	if !helpers.CheckError(ctx, err) {
		return
	}

	token, err := helpers.GenerateToken(user.ID, user.Name, user.Email, string(user.Role))
	if !helpers.CheckError(ctx, err) {
		return
	}

	result := dto.LoginResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Role:  string(user.Role),
		Token: token,
	}

	helpers.ResponseSuccess(ctx, http.StatusOK, "login sukses", result)
}

func (c *UserController) GetProfile(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if !helpers.CheckError(ctx, err) {
		return
	}

	user, err := c.service.GetProfile(uint(id))
	if !helpers.CheckError(ctx, err) {
		return
	}

	result := dto.UpdateProfileResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}

	helpers.ResponseSuccess(ctx, http.StatusOK, "success", result)
}

func (c *UserController) UpdateProfile(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if !helpers.CheckError(ctx, err) {
		return
	}

	var input dto.UpdateProfileRequest
	if !helpers.BindAndValidate(ctx, &input) {
		return
	}

	updated, err := c.service.UpdateProfile(uint(id), input.Name, input.Email)
	if !helpers.CheckError(ctx, err) {
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
	id, err := strconv.Atoi(idParam)
	if !helpers.CheckError(ctx, err) {
		return
	}

	err = c.service.DeleteProfile(uint(id))
	if !helpers.CheckError(ctx, err) {
		return
	}

	helpers.ResponseSuccess(ctx, http.StatusOK, "success delete user", nil)
}
