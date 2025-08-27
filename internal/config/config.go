package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ENVIRONMENT          string
	DB_URI               string
	GRPC_PORT            string
	REST_PORT            string
	JWT_SECRET           string
	CORS_ALLOWED_ORIGINS string
	STORAGE_SERVICE_URL  string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()

	if err != nil {
		log.Println("Warning: Could not load .env file. Falling back to environment variables.")
	} else {
		log.Println("Config loaded from .env file")
	}

	return &Config{
		ENVIRONMENT:          os.Getenv("ENVIRONMENT"),
		DB_URI:               os.Getenv("DB_URI"),
		GRPC_PORT:            os.Getenv("GRPC_PORT"),
		REST_PORT:            os.Getenv("REST_PORT"),
		JWT_SECRET:           os.Getenv("JWT_SECRET"),
		CORS_ALLOWED_ORIGINS: os.Getenv("CORS_ALLOWED_ORIGINS"),
		STORAGE_SERVICE_URL:  os.Getenv("STORAGE_SERVICE_URL"),
	}, nil

}
