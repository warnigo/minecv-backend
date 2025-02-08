package resume_templates

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"minecv/internal/domain/entities"
	"minecv/internal/domain/schemas"
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
	utils.RespondSuccess(c, http.StatusOK, templates, translate("templates.success", nil))
}

// GetTemplateByID handles GET /templates/:id - Fetch a single template by ID
func (controllers *TemplateController) GetTemplateByID(c *gin.Context) {
	localizer, _ := c.Get(("localizer"))
	i18n := localizer.(*localization.I18n)
	lang := c.GetString(("lang"))
	translate := utils.GetTranslator(i18n, lang)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.RespondError(c, http.StatusBadRequest, translate("validations.invalid_id", nil))
		return
	}
	template, err := controllers.Service.GetTemplateByID(uint(id))
	if err != nil {
		utils.RespondError(c, http.StatusNotFound, translate("templates.not_found", nil))
		return
	}
	utils.RespondSuccess(c, http.StatusOK, template, translate("templates.success", nil))
}

// CreateTemplate handles POST /templates - Create a new template
func (controllers *TemplateController) CreateTemplate(c *gin.Context) {
	localizer, _ := c.Get("localizer")
	i18n := localizer.(*localization.I18n)
	lang := c.GetString("lang")
	translate := utils.GetTranslator(i18n, lang)

	var req schemas.CreateResumeTemplateSchemas
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.RespondError(c, http.StatusBadRequest, translate("validations.invalid_input", nil))
		return
	}

	template := entities.ResumeTemplateEntity{
		Title:              req.Title,
		Description:        req.Description,
		Category:           req.Category,
		HTMLContent:        req.HTMLContent,
		CSSContent:         req.CSSContent,
		CSSVariables:       req.CSSVariables,
		Sections:           req.Sections,
		ThumbnailImg:       req.ThumbnailImg,
		PreviewURL:         req.PreviewURL,
		DefaultFont:        req.DefaultFont,
		DefaultColorScheme: req.DefaultColorScheme,
		Compatibility:      entities.PaperFormat(req.Compatibility),
	}

	createdTemplate, err := controllers.Service.CreateTemplate(template)
	if err != nil {
		utils.RespondError(c, http.StatusInternalServerError, translate("templates.failedCreateTemplate", nil))
		return
	}

	utils.RespondSuccess(c, http.StatusCreated, createdTemplate, translate("templates.createdSuccessfully", nil))
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
