package repositories

import (
	"minecv/internal/domain/entities"
	"minecv/internal/infrastructure/database"
)

// TemplateRepository provides methods to interact with the database
type TemplateRepository struct{}

// GetAll retrieves all resume templates
func (r *TemplateRepository) GetAll() ([]entities.ResumeTemplateEntity, error) {
	var templates []entities.ResumeTemplateEntity
	err := database.DB.Find(&templates).Error
	return templates, err
}

// GetByID retrieves a resume resume_templates by ID
func (r *TemplateRepository) GetByID(id uint) (entities.ResumeTemplateEntity, error) {
	var template entities.ResumeTemplateEntity
	err := database.DB.First(&template, id).Error
	return template, err
}

// Create adds a new resume resume_templates to the database
func (r *TemplateRepository) Create(template entities.ResumeTemplateEntity) error {
	return database.DB.Create(&template).Error
}

// Update modifies an existing resume resume_templates
func (r *TemplateRepository) Update(template entities.ResumeTemplateEntity) error {
	return database.DB.Save(&template).Error
}

// Delete removes a resume resume_templates by ID
func (r *TemplateRepository) Delete(id uint) error {
	return database.DB.Delete(&entities.ResumeTemplateEntity{}, id).Error
}
