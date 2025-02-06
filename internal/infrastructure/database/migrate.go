package database

import (
	"log"

	"minecv/internal/domain/entities"
	"minecv/pkg/utils"
)

// AutoMigrate perform the database migrations for all modals
func AutoMigrate() {
	err := DB.AutoMigrate(
		&entities.UserEntity{},
		&entities.ResumeTemplateEntity{},
	)
	if err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}
	utils.LogSuccess("Database migrations completed successfully!")
}
