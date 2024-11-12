package config

import (
	"log"

	"github.com/joho/godotenv"
)

// LoadConfig loads environment variables
func LoadConfig() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
}
