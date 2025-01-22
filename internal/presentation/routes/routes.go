package routes

import (
	"github.com/gin-gonic/gin"

	"minecv/internal/presentation/controllers"
)

// RegisterRoutes sets up the application routes
func RegisterRoutes(public *gin.RouterGroup, protected *gin.RouterGroup) {
	// Public routes
	authGroup := public.Group("/auth")
	{
		authGroup.POST("/register", controllers.Register)
		authGroup.POST("/login", controllers.Login)
		authGroup.POST("/refresh", controllers.RefreshToken)
	}

	// Protected routes
}
