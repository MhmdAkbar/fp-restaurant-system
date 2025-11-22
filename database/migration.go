package database

import (
	"aplikasi_restoran/internal/models"
	"log" // log error
	"gorm.io/gorm" // tipe *gorm.DB
)

func RunMigration(db *gorm.DB) { // fungsi menerima koneksi DB
	err := db.AutoMigrate(&models.User{}) // migrasi tabel user
	if err != nil {
		log.Fatal("gagal migrate user:", err) // log error kalau gagal
	}
}
