package bootstrap

import (
	"gorm.io/gorm"

	tablecontroller "aplikasi_restoran/internal/controllers/table"
	usercontroller "aplikasi_restoran/internal/controllers/user"
	tablerepository "aplikasi_restoran/internal/repositories/table"
	userrepo "aplikasi_restoran/internal/repositories/user"
	tableservice "aplikasi_restoran/internal/services/table"
	userservice "aplikasi_restoran/internal/services/user"
)

type AppModule struct {
	UserController *usercontroller.UserController
	TableController *tablecontroller.TableController
}

func InitModules(db *gorm.DB) *AppModule {
	userRepo := userrepo.NewUserRepository(db)
	userService := userservice.NewUserService(userRepo)
	userController := usercontroller.NewController(userService)

	tableRepo := tablerepository.NewTableRepository(db)
	tableService := tableservice.NewTableService(tableRepo)
	tableController := tablecontroller.NewController(tableService)

	return &AppModule{
		UserController: userController,
		TableController: tableController,
	}
}
