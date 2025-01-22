package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"minecv/config"
)

var DB *gorm.DB

// InitDatabase initialize the database connection
func InitDatabase() {
	// Constructor the DSN for postgres connection
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		config.AppConfig.PostgresHost,
		config.AppConfig.PostgresUser,
		config.AppConfig.PostgresPassword,
		config.AppConfig.PostgresDatabase,
		config.AppConfig.PostgresPort,
	)

	// Connection to the database
	var dbErr error
	DB, dbErr = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if dbErr != nil {
		log.Printf("DSN: %s", dsn)
		log.Fatalf("Failed to connect to the database: %v", dbErr)
	}

	log.Printf("Successfully connected to the database")
}
