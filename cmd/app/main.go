package main

import (
	"aplikasi_restoran/config"
	"aplikasi_restoran/database"
	bootstrap "aplikasi_restoran/init"
	"aplikasi_restoran/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.Init()
	db := database.NewMySqlConnection(config.GetDSN())
	database.RunMigration(db)

	modules := bootstrap.InitModules(db)

	r := gin.Default()
	routes.UserRoutes(r, modules.UserController)
	routes.TableRouter(r, modules.TableController, modules.OrderController)
	routes.MenuRouter(r, modules.MenuController)
	routes.OrderRouter(r, modules.OrderController, modules.OrderDetailController)

	r.Run(config.GetAppPort())
}
