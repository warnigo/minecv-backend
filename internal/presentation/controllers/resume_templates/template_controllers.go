package resume_templates

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"minecv/internal/domain/entities"
	"minecv/internal/domain/services"
	"minecv/internal/infrastructure/localization"
	"minecv/pkg/utils"
)

// TemplateController handles HTTP requests related to resume templates
type TemplateController struct {
	Service services.TemplateService
}

// GetTemplates handles GET /templates - Fetch all templates
func (controllers *TemplateController) GetTemplates(c *gin.Context) {
	localizer, _ := c.Get("localizer")
	i18n := localizer.(*localization.I18n)
	lang := c.GetString("lang")
	translate := utils.GetTranslator(i18n, lang)

	templates, err := controllers.Service.GetAllTemplates()
	if err != nil {
		utils.RespondError(c, http.StatusInternalServerError, translate("templates.failed_fetch", nil))
		return
	}
	c.JSON(http.StatusOK, templates)
}

// GetTemplateByID handles GET /templates/:id - Fetch a single template by ID
func (controllers *TemplateController) GetTemplateByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	template, err := controllers.Service.GetTemplateByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Template not found"})
		return
	}
	c.JSON(http.StatusOK, template)
}

// CreateTemplate handles POST /templates - Create a new template
func (controllers *TemplateController) CreateTemplate(c *gin.Context) {
	var template entities.ResumeTemplateEntity
	if err := c.ShouldBindJSON(&template); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	err := controllers.Service.CreateTemplate(template)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create template"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Template created successfully"})
}

// UpdateTemplate handles PUT /templates/:id - Update an existing template
func (controllers *TemplateController) UpdateTemplate(c *gin.Context) {
	var template entities.ResumeTemplateEntity
	if err := c.ShouldBindJSON(&template); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	err := controllers.Service.UpdateTemplate(template)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update template"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Template updated successfully"})
}

// DeleteTemplate handles DELETE /templates/:id - Delete a template
func (controllers *TemplateController) DeleteTemplate(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	err = controllers.Service.DeleteTemplate(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete template"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Template deleted successfully"})
}
