package config

import (
	"log" // log pesan error
	"os"  // ambil variabel env OS

	"github.com/joho/godotenv" // load file .env
)

func LoadEnv() {
	err := godotenv.Load("./cmd/app/.env") // load .env dari path pasti
	if err != nil {
		log.Println("file .env tidak ditemukan:", err) // tampilkan error asli
	}
}


func GetDSN() string { // ambil DSN string
	return os.Getenv("DB_DSN") // ambil DB_DSN dari .env
}
