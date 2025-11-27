package config

import (
	"log"
	"github.com/joho/godotenv"
)

func Init() {
	err := godotenv.Load("./cmd/app/.env") // otomatis cari file .env di root project
	if err != nil {
		log.Println("⚠️ .env not found, using system env...")
	}
}
