package config

import (
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	_ = godotenv.Load()
}

func GetPort() string {

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	return port
}
