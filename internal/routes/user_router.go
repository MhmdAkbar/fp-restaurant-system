package routes

import (
	usercontroller "aplikasi_restoran/internal/controllers/user"
	"aplikasi_restoran/internal/middlewares"
	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine, uc *usercontroller.UserController) {
	users := r.Group("/users")

	// Public
	users.POST("/login", uc.Login)

	// Authorized user routes
	usersAuth := users.Group("")
	usersAuth.Use(middlewares.AuthMiddleware())
	{
		usersAuth.GET("/:id", middlewares.VerifyUserAccess(), uc.GetProfile)
		usersAuth.PATCH("/:id", middlewares.VerifyUserAccess(), uc.UpdateProfile)
	}

	// Super admin only
	admin := users.Group("")
	admin.Use(middlewares.AuthMiddleware(), middlewares.Role("super_admin"))
	{
		admin.POST("", uc.Register)     // Create new user
		admin.DELETE("/:id", uc.DeleteProfile)
	}
}
