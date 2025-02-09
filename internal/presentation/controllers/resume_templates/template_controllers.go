package resume_templates

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/datatypes"
	"minecv/internal/domain/schemas"
	"net/http"

	"minecv/internal/domain/entities"
	"minecv/internal/domain/services"
	"minecv/pkg/utils"
)

// TemplateController handles HTTP requests related to resume templates
type TemplateController struct {
	Service services.TemplateService
}

// GetTemplates handles GET /templates - Fetch all templates
func (controllers *TemplateController) GetTemplates(c *gin.Context) {
	translate := c.MustGet("translator").(utils.TranslatorFunc)

	templates, err := controllers.Service.GetAllTemplates()
	if err != nil {
		utils.RespondError(c, http.StatusInternalServerError, translate("templates.failed_fetch"))
		return
	}
	utils.RespondSuccess(c, http.StatusOK, templates, translate("templates.success"))
}

// GetTemplateByID handles GET /templates/:id - Fetch a single template by ID
func (controllers *TemplateController) GetTemplateByID(c *gin.Context) {
	translate := c.MustGet("translator").(utils.TranslatorFunc)

	id := c.Param("id")
	if id == "" {
		utils.RespondError(c, http.StatusBadRequest, translate("validations.invalid_id"))
		return
	}
	template, err := controllers.Service.GetTemplateByID(id)
	if err != nil {
		utils.RespondError(c, http.StatusNotFound, translate("templates.not_found"))
		return
	}
	utils.RespondSuccess(c, http.StatusOK, template, translate("templates.success"))
}

// CreateTemplate handles POST /templates - Create a new template
func (controllers *TemplateController) CreateTemplate(c *gin.Context) {
	var template schemas.CreateResumeTemplateSchemas

	translate := c.MustGet("translator").(utils.TranslatorFunc)

	if err := c.ShouldBindJSON(&template); err != nil {
		utils.RespondError(c, http.StatusBadRequest, translate("validations.invalid_input"))
		fmt.Println(err)
		return
	}

	// Convert schema to entity
	resumeTemplate := entities.ResumeTemplateEntity{
		ID:                 uuid.New().String(),
		Title:              template.Title,
		Description:        template.Description,
		Category:           template.Category,
		Language:           template.Language,
		Tags:               datatypes.JSON(utils.ToJSON(template.Tags)),
		HTMLContent:        template.HTMLContent,
		CSSContent:         template.CSSContent,
		Sections:           datatypes.JSON(utils.ToJSON(template.Sections)),
		ThumbnailImg:       template.ThumbnailImg,
		PreviewURL:         template.PreviewURL,
		DefaultFont:        template.DefaultFont,
		DefaultColorScheme: template.DefaultColorScheme,
		Compatibility:      entities.PaperFormat(template.Compatibility),
		Downloads:          template.Downloads,
		IsActive:           template.IsActive,
		IsFeatured:         template.IsFeatured,
		Rating:             template.Rating,
		License:            template.License,
		IsPremium:          template.IsPremium,
	}

	createdTemplate, err := controllers.Service.CreateTemplate(resumeTemplate)
	if err != nil {
		utils.RespondError(c, http.StatusInternalServerError, translate("templates.failedCreateTemplate"))
		return
	}

	utils.RespondSuccess(c, http.StatusCreated, createdTemplate, translate("templates.createdSuccessfully"))
}

// UpdateTemplate handles PUT /templates/:id - Update an existing template
func (controllers *TemplateController) UpdateTemplate(c *gin.Context) {
	translate := c.MustGet("translator").(utils.TranslatorFunc)

	id := c.Param("id")
	if id == "" {
		utils.RespondError(c, http.StatusBadRequest, translate("validations.invalid_id"))
		return
	}

	existingTemplate, err := controllers.Service.GetTemplateByID(id)
	if err != nil {
		utils.RespondError(c, http.StatusNotFound, translate("templates.not_found"))
		return
	}

	var updatedData schemas.CreateResumeTemplateSchemas
	if err := c.ShouldBindJSON(&updatedData); err != nil {
		utils.RespondError(c, http.StatusBadRequest, translate("validations.invalid_input"))
		return
	}

	existingTemplate.Title = updatedData.Title
	existingTemplate.Description = updatedData.Description
	existingTemplate.Category = updatedData.Category
	existingTemplate.Language = updatedData.Language
	existingTemplate.Tags = datatypes.JSON(utils.ToJSON(updatedData.Tags))
	existingTemplate.HTMLContent = updatedData.HTMLContent
	existingTemplate.CSSContent = updatedData.CSSContent
	existingTemplate.Sections = datatypes.JSON(utils.ToJSON(updatedData.Sections))
	existingTemplate.ThumbnailImg = updatedData.ThumbnailImg
	existingTemplate.PreviewURL = updatedData.PreviewURL
	existingTemplate.DefaultFont = updatedData.DefaultFont
	existingTemplate.DefaultColorScheme = updatedData.DefaultColorScheme
	existingTemplate.Compatibility = entities.PaperFormat(updatedData.Compatibility)
	existingTemplate.Downloads = updatedData.Downloads
	existingTemplate.IsActive = updatedData.IsActive
	existingTemplate.IsFeatured = updatedData.IsFeatured
	existingTemplate.Rating = updatedData.Rating
	existingTemplate.License = updatedData.License
	existingTemplate.IsPremium = updatedData.IsPremium

	updatedTemplate, err := controllers.Service.UpdateTemplate(existingTemplate)
	if err != nil {
		utils.RespondError(c, http.StatusInternalServerError, translate("templates.failed_update"))
		return
	}

	utils.RespondSuccess(c, http.StatusOK, updatedTemplate, translate("templates.update_successfully"))
}

// DeleteTemplate handles DELETE /templates/:id - Delete a template
func (controllers *TemplateController) DeleteTemplate(c *gin.Context) {
	translate := c.MustGet("translator").(utils.TranslatorFunc)

	id := c.Param("id")
	if id == "" {
		utils.RespondError(c, http.StatusBadRequest, translate("validations.invalid_id"))
		return
	}

	err := controllers.Service.DeleteTemplate(id)
	if err != nil {
		utils.RespondError(c, http.StatusInternalServerError, translate("templates.delete_failed"))
		return
	}

	utils.RespondSuccess(c, http.StatusOK, nil, translate("templates.deleted_successfully"))
}
