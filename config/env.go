package config

import "os"

func GetDSN() string {
	return os.Getenv("DB_DSN")
}

func GetJWTSecret() []byte {
	return []byte(os.Getenv("JWT_SECRET"))
}

func GetAppPort() string {
	return os.Getenv("APP_PORT")
}
