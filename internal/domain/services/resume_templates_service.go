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
func (s *TemplateService) GetTemplateByID(id string) (entities.ResumeTemplateEntity, error) {
	return s.Repo.GetByID(id)
}

// CreateTemplate creates a new resume resume_templates
func (s *TemplateService) CreateTemplate(template entities.ResumeTemplateEntity) (entities.ResumeTemplateEntity, error) {
	createdTemplate, err := s.Repo.Create(template)
	if err != nil {
		return entities.ResumeTemplateEntity{}, err
	}
	return createdTemplate, nil
}

// UpdateTemplate updates an existing resume resume_templates
func (s *TemplateService) UpdateTemplate(template entities.ResumeTemplateEntity) (entities.ResumeTemplateEntity, error) {
	updatedTemplate, err := s.Repo.Update(template)
	if err != nil {
		return entities.ResumeTemplateEntity{}, err
	}
	return updatedTemplate, nil
}

// DeleteTemplate deletes a resume resume_templates by ID
func (s *TemplateService) DeleteTemplate(id string) error {
	return s.Repo.Delete(id)
}
