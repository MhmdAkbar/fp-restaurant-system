package routes

import (
	usercontroller "aplikasi_restoran/internal/controllers/user"
	"aplikasi_restoran/internal/middlewares"

	"github.com/gin-gonic/gin" // framework gin
)

func UserRoutes(r *gin.Engine, uc *usercontroller.Controller) {
	user := r.Group("/users")
	{
		user.POST("/login", uc.Login)

		// Protected routes (hanya pemilik akun atau super_admin)
		protected := user.Group("/")
		protected.Use(middlewares.AuthMiddleware())
		{
			protected.GET("/:id", middlewares.VerifyUserAccess(), uc.GetProfile)
			protected.PUT("/:id/update", middlewares.VerifyUserAccess(), uc.UpdateProfile)
		}

		// Super admin only
		admin := user.Group("/")
		admin.Use(middlewares.AuthMiddleware(), middlewares.Role("super_admin"))
		{
			admin.DELETE("/:id/delete", uc.DeleteProfile)
			admin.POST("/register", uc.Register)
		}
	}
}
