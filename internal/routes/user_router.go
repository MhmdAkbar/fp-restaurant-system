package routes

import (
	usercontroller "aplikasi_restoran/internal/controllers/user"

	"github.com/gin-gonic/gin" // framework gin
)

func UserRoutes(r *gin.Engine, uc *usercontroller.Controller) { // fungsi definisi route
	user := r.Group("/users") // grup route /users
	{
		user.POST("/register", uc.Register) // route register
		user.POST("/login", uc.Login)       // route login
		user.GET("/:id", uc.GetProfile)
	}
}
