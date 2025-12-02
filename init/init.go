package bootstrap

import (
	"gorm.io/gorm"

	menucontroller "aplikasi_restoran/internal/controllers/menu"
	tablecontroller "aplikasi_restoran/internal/controllers/table"
	usercontroller "aplikasi_restoran/internal/controllers/user"
	menurepository "aplikasi_restoran/internal/repositories/menu"
	tablerepository "aplikasi_restoran/internal/repositories/table"
	userrepo "aplikasi_restoran/internal/repositories/user"
	menuservice "aplikasi_restoran/internal/services/menu"
	tableservice "aplikasi_restoran/internal/services/table"
	userservice "aplikasi_restoran/internal/services/user"
)

type AppModule struct {
	UserController *usercontroller.UserController
	TableController *tablecontroller.TableController
	MenuController *menucontroller.MenuController
}

func InitModules(db *gorm.DB) *AppModule {
	userRepo := userrepo.NewUserRepository(db)
	userService := userservice.NewUserService(userRepo)
	userController := usercontroller.NewController(userService)

	tableRepo := tablerepository.NewTableRepository(db)
	tableService := tableservice.NewTableService(tableRepo)
	tableController := tablecontroller.NewController(tableService)

	menuRepo := menurepository.NewMenuRepository(db)
	menuService := menuservice.NewMenuService(menuRepo)
	menuController := menucontroller.NewController(menuService)
	return &AppModule{
		UserController: userController,
		TableController: tableController,
		MenuController: menuController,
	}
}
