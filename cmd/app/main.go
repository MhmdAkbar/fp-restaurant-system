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
	config.LoadEnv()
	db := database.NewMySqlConnection(config.GetDSN()) // now returns *gorm.DB
	database.RunMigration(db)

	repo := repoitories.NewUserRepository(db) // gunakan koneksi nyata
	service := userservice.NewUserService(repo)
	controller := usercontroller.New(service) // panggil konstruktor yang benar

	router := gin.Default()

	routes.UserRoutes(router, controller)

	router.Run(":8080")
}
