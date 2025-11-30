package database

import (
	tablesmodels "aplikasi_restoran/internal/models/tables"
	usermodels "aplikasi_restoran/internal/models/user"
	"log"

	"gorm.io/gorm"
)

func RunMigration(db *gorm.DB) { // fungsi menerima koneksi DB
	err := db.AutoMigrate(&usermodels.User{}, &tablesmodels.Table{}) 
	if err != nil {
		log.Fatal("gagal migrate user:", err) // log error kalau gagal
	}
}
