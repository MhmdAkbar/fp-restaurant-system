package database

import (
	// core gorm
	"log" // untuk log error

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMySqlConnection(dsn string) *gorm.DB {

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{}) // buka koneksi GORM ke MySQL
	if err != nil {
		log.Fatal("gagal connect ke database:", err) // log error kalau gagal connect
	}

	sqlDB, err := db.DB() // dapatkan objek sql.DB bawaan Go
	if err != nil {
		log.Fatal("gagal ambil sql.DB:", err) // handle kalau ada error
	}

	err = sqlDB.Ping() // tes koneksi berjalan atau tidak
	if err != nil {
		log.Fatal("ping gagal:", err) // gagal ping → database tidak reachable
	}

	log.Println("Koneksi ke database berhasil!") 
	return db
}
