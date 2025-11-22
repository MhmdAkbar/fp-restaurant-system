package usercontroller

import (
	"aplikasi_restoran/internal/models"
	userservice "aplikasi_restoran/internal/services/user"
	"fmt"
	"net/http" // http response

	"github.com/gin-gonic/gin" // framework gin
)

type Controller struct { // controller user
	service userservice.UserService
}

func New(service userservice.UserService) *Controller {
	return &Controller{service} // buat instance
}

func (c *Controller) Register(ctx *gin.Context) { // handler register
	var input struct { // input JSON
		Name     string `json:"name"`     // nama
		Email    string `json:"email"`    // email
		Password string `json:"password"` // password
		Role     string `json:"role"`     // role
	}

	if err := ctx.ShouldBindJSON(&input); err != nil { // parsing JSON
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) // error parsing
		return
	}

	err := c.service.Register(input.Name, input.Email, input.Password, models.UserRole(input.Role)) // panggil service
	if err != nil {                                                                                 // cek error service
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) // error service
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "register sukses"}) // respon sukses
}

func (c *Controller) Login(ctx *gin.Context) { // handler login
	var input struct { // input JSON
		Email    string `json:"email"`    // email
		Password string `json:"password"` // password
	}

	if err := ctx.ShouldBindJSON(&input); err != nil { // parsing JSON
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) // error
		return
	}

	user, err := c.service.Login(input.Email, input.Password) // panggil service
	if err != nil {                                           // login gagal
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()}) // 401
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"user": user}) // respon user
}

func (c *Controller) GetProfile(ctx *gin.Context) { // handler get user by id
	idParam := ctx.Param("id") // ambil parameter id dari URL
	if idParam == "" {         // cek kosong
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "id wajib diisi"}) // respon error
		return                                                            // stop eksekusi
	}

	var id uint                       // siapkan variabel id uint
	_, err := fmt.Sscan(idParam, &id) // convert string ke uint
	if err != nil {                   // cek salah konversi
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "id tidak valid"}) // error
		return                                                            // stop
	}

	user, err := c.service.GetProfile(id) // panggil service
	if err != nil {                       // cek jika tidak ditemukan
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()}) // error 404
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"user": user}) // respon sukses
}
