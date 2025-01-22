package main

import (
	"minecv/config"
	"minecv/internal/infrastructure/database"
	"minecv/internal/infrastructure/http"
)

// @title Minecv backend
// @description This platform is completely free to use and enables anyone to build high-quality resumes.
// @version: 1.0
// @BasePath /
func main() {
	// Load the configuration from environment variables
	config.LoadConfig()

	// Initialize the database
	database.InitDatabase()
	// AutoMigrate the database column
	database.AutoMigrate()

	// Initialize routes with gin
	http.InitialServer()
}
