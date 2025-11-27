package main

import (
	"aplikasi_restoran/config"
	"aplikasi_restoran/database"
	usercontroller "aplikasi_restoran/internal/controllers/user"
	repoitories "aplikasi_restoran/internal/repositories/user"
	"aplikasi_restoran/internal/routes"
	userservice "aplikasi_restoran/internal/services/user"

	"github.com/gin-gonic/gin"
)

func main() {
	config.Init()
	db := database.NewMySqlConnection(config.GetDSN()) // mengambil string dsn dari .env
	database.RunMigration(db)

	repo := repoitories.NewUserRepository(db)
	service := userservice.NewUserService(repo)
	controller := usercontroller.NewController(service) // panggil konstruktor yang benar

	router := gin.Default()

	routes.UserRoutes(router, controller)
	port := config.GetAppPort()
	router.Run(port)
}
