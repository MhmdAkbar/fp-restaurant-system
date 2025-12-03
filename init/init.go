package bootstrap

import (
	"gorm.io/gorm"

	menucontroller "aplikasi_restoran/internal/controllers/menu"
	ordercontroller "aplikasi_restoran/internal/controllers/order"
	orderdetailcontroller "aplikasi_restoran/internal/controllers/order_detail"
	tablecontroller "aplikasi_restoran/internal/controllers/table"
	usercontroller "aplikasi_restoran/internal/controllers/user"
	menurepository "aplikasi_restoran/internal/repositories/menu"
	orderrepository "aplikasi_restoran/internal/repositories/order"
	orderdetailrepository "aplikasi_restoran/internal/repositories/order_detail"
	tablerepository "aplikasi_restoran/internal/repositories/table"
	userrepo "aplikasi_restoran/internal/repositories/user"
	menuservice "aplikasi_restoran/internal/services/menu"
	orderservice "aplikasi_restoran/internal/services/order"
	orderdetailservice "aplikasi_restoran/internal/services/order_detail"
	tableservice "aplikasi_restoran/internal/services/table"
	userservice "aplikasi_restoran/internal/services/user"
)

type AppModule struct {
	UserController *usercontroller.UserController
	TableController *tablecontroller.TableController
	MenuController *menucontroller.MenuController
	OrderController *ordercontroller.OrderController
	OrderDetailController *orderdetailcontroller.OrderDetailController
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
orderRepo := orderrepository.NewOrderRepository(db)
orderDetailRepo := orderdetailrepository.NewOrderDetailRepository(db)

orderService := orderservice.NewOrderService(orderRepo, orderDetailRepo, menuRepo)
orderController := ordercontroller.NewOrderController(orderService)

orderDetailService := orderdetailservice.NewOrderDetailService(orderDetailRepo, menuRepo, orderRepo,orderService)
orderDetailController := orderdetailcontroller.NewOrderDetailController(orderDetailService)

	return &AppModule{
		UserController: userController,
		TableController: tableController,
		MenuController: menuController,
		OrderController: orderController,
		OrderDetailController: orderDetailController,
	}
}
