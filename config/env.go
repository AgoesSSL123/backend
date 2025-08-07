package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, relying on system environment variables.")
	}
}

func GetPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // default port
	}
	return port
}
