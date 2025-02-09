package routes

import (
	"github.com/gin-gonic/gin"
	"minecv/internal/presentation/controllers/auth"
	"minecv/internal/presentation/controllers/resume_templates"
)

func RegisterRoutes(public *gin.RouterGroup, protected *gin.RouterGroup) {
	ctrl := resume_templates.TemplateController{}

	// Public routes
	authGroup := public.Group("/app/auth")
	{
		authGroup.POST("/register", auth.Register)
		authGroup.POST("/login", auth.Login)
		authGroup.POST("/refresh", auth.RefreshToken)
	}
	templates := public.Group("/resume/templates")
	{
		templates.GET("/", ctrl.GetTemplates)
		templates.GET("/:id", ctrl.GetTemplateByID)
		templates.POST("/", ctrl.CreateTemplate)
		templates.PUT("/", ctrl.UpdateTemplate)
		templates.DELETE("/:id", ctrl.DeleteTemplate)
	}

	// Protected routes
}
