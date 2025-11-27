package database

import (
	usermodels "aplikasi_restoran/internal/models/user"
	"log"
	"gorm.io/gorm"
)

func RunMigration(db *gorm.DB) { // fungsi menerima koneksi DB
	err := db.AutoMigrate(&usermodels.User{}) // migrasi tabel user
	if err != nil {
		log.Fatal("gagal migrate user:", err) // log error kalau gagal
	}
}
