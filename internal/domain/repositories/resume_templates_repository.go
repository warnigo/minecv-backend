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
func (r *TemplateRepository) GetByID(id string) (entities.ResumeTemplateEntity, error) {
	var template entities.ResumeTemplateEntity
	err := database.DB.Where("id = ?", id).First(&template).Error
	return template, err
}

// Create adds a new resume resume_templates to the database
func (r *TemplateRepository) Create(template entities.ResumeTemplateEntity) (entities.ResumeTemplateEntity, error) {
	if err := database.DB.Create(&template).Error; err != nil {
		return entities.ResumeTemplateEntity{}, err
	}
	return template, nil
}

// Update modifies an existing resume template and returns the updated entity
func (r *TemplateRepository) Update(template entities.ResumeTemplateEntity) (entities.ResumeTemplateEntity, error) {
	if err := database.DB.Save(&template).Error; err != nil {
		return entities.ResumeTemplateEntity{}, err
	}
	return template, nil
}

// Delete removes a resume resume_templates by ID
func (r *TemplateRepository) Delete(id string) error {
	return database.DB.Where("id = ?", id).Delete(&entities.ResumeTemplateEntity{}).Error
}
