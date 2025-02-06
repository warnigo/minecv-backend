package services

import (
	"minecv/internal/domain/entities"
	"minecv/internal/domain/repositories"
)

// TemplateService handles business logic for resume templates
type TemplateService struct {
	Repo repositories.TemplateRepository
}

// GetAllTemplates returns all resume templates
func (s *TemplateService) GetAllTemplates() ([]entities.ResumeTemplateEntity, error) {
	return s.Repo.GetAll()
}

// GetTemplateByID returns a resume resume_templates by ID
func (s *TemplateService) GetTemplateByID(id uint) (entities.ResumeTemplateEntity, error) {
	return s.Repo.GetByID(id)
}

// CreateTemplate creates a new resume resume_templates
func (s *TemplateService) CreateTemplate(template entities.ResumeTemplateEntity) error {
	return s.Repo.Create(template)
}

// UpdateTemplate updates an existing resume resume_templates
func (s *TemplateService) UpdateTemplate(template entities.ResumeTemplateEntity) error {
	return s.Repo.Update(template)
}

// DeleteTemplate deletes a resume resume_templates by ID
func (s *TemplateService) DeleteTemplate(id uint) error {
	return s.Repo.Delete(id)
}
