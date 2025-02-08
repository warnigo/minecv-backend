package routes

import (
	"github.com/gin-gonic/gin"
	"minecv/internal/presentation/controllers/auth"
	"minecv/internal/presentation/controllers/resume_templates"
)

func RegisterRoutes(public *gin.RouterGroup, protected *gin.RouterGroup) {
	// Public routes
	authGroup := public.Group("/auth")
	{
		authGroup.POST("/register", auth.Register)
		authGroup.POST("/login", auth.Login)
		authGroup.POST("/refresh", auth.RefreshToken)
	}

	// Protected routes
	ctrl := resume_templates.TemplateController{}
	templates := protected.Group("/resume/templates")
	{
		templates.GET("/", ctrl.GetTemplates)
		templates.GET("/:id", ctrl.GetTemplateByID)
		templates.POST("/", ctrl.CreateTemplate)
		templates.PUT("/", ctrl.UpdateTemplate)
		templates.DELETE("/:id", ctrl.DeleteTemplate)
	}
}
