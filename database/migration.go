package database

import (
	"aplikasi_restoran/internal/models"
	"log"
	"gorm.io/gorm"
)

func RunMigration(db *gorm.DB) { // fungsi menerima koneksi DB
	err := db.AutoMigrate(&models.User{}, &models.Table{}, &models.Menu{})
	if err != nil {
		log.Fatal("gagal migrate user:", err) // log error kalau gagal
	}
}
