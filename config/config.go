package config

import (
	"log"
	"os"
	
	"github.com/joho/godotenv"
)

type Config struct {
	ApplicationVariable string
	isDevelopment       bool

	PostgresHost     string
	PostgresPort     string
	PostgresUser     string
	PostgresPassword string
	PostgresDatabase string

	ApiVersion string
	ApiPort    string
}

// AppConfig holds the global configuration accessible by the entire application
var AppConfig Config

// LoadConfig leads environment variables into to AppConfig struct
func LoadConfig() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Load configuration from environment variables
	AppConfig.ApplicationVariable = os.Getenv("APPLICATION_ENVIRONMENT")
	AppConfig.isDevelopment = AppConfig.ApplicationVariable == "development"

	AppConfig.ApiVersion = os.Getenv("API_VERSION")
	AppConfig.ApiPort = os.Getenv("API_PORT")

	AppConfig.PostgresHost = os.Getenv("POSTGRES_HOST")
	AppConfig.PostgresPort = os.Getenv("POSTGRES_PORT")
	AppConfig.PostgresUser = os.Getenv("POSTGRES_USER")
	AppConfig.PostgresPassword = os.Getenv("POSTGRES_PASSWORD")
	AppConfig.PostgresDatabase = os.Getenv("POSTGRES_DATABASE")

	if AppConfig.ApiVersion == "" || AppConfig.ApiPort == "" {
		log.Fatal("Some required environment variables are missing, port and versions!")
	}

	if AppConfig.PostgresHost == "" || AppConfig.PostgresPort == "" {
		log.Fatal("Some required environment variables are missing!")
	}
}
