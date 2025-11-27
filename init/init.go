package bootstrap

import (
	"gorm.io/gorm"

	userrepo "aplikasi_restoran/internal/repositories/user"
	userservice "aplikasi_restoran/internal/services/user"
	usercontroller "aplikasi_restoran/internal/controllers/user"
)

type AppModule struct {
	UserController *usercontroller.Controller
}

func InitModules(db *gorm.DB) *AppModule {
	userRepo := userrepo.NewUserRepository(db)
	userService := userservice.NewUserService(userRepo)
	userController := usercontroller.NewController(userService)

	return &AppModule{
		UserController: userController,
	}
}
