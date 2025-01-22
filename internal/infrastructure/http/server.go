package http

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"minecv/config"
	"minecv/internal/infrastructure/localization"
	"minecv/internal/presentation/middleware"
	"minecv/internal/presentation/routes"
	"minecv/pkg/utils"
)

// InitialServer initializes and returns a Gin routes instance
func InitialServer() {
	// Set Gin mode based on environment
	switch config.AppConfig.ApplicationVariable {
	case "development":
		gin.SetMode(gin.DebugMode)
	case "production":
		gin.SetMode(gin.ReleaseMode)
	case "testing":
		gin.SetMode(gin.TestMode)
	default:
		gin.SetMode(gin.ReleaseMode)
	}

		// Initialize I18n
	i18n := localization.NewI18n("en")
	if i18n == nil {
		fmt.Println("Failed to initialize I18n")
		return
	}

	// Create Gin instance
	router := gin.Default()

	// Apply LocalizationMiddleware globally
	router.Use(middleware.LocalizationMiddleware(i18n))

	// Define API version group
	apiVersion := fmt.Sprintf("/api/%s", config.AppConfig.ApiVersion)
	apiGroup := router.Group(apiVersion)

	// Define Public and Protected Groups
	publicGroup := apiGroup
	protectedGroup := apiGroup.Group("/app")

	// Add middleware to handle invalid methods and routes
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"error":  "Route not found",
			"method": c.Request.Method,
			"path":   c.Request.URL.Path,
		})
	})
	
	router.NoMethod(func(c *gin.Context) {
		c.JSON(http.StatusMethodNotAllowed, gin.H{
			"error":  "Method not allowed",
			"method": c.Request.Method,
			"path":   c.Request.URL.Path,
		})
	})

	// Register routes
	routes.RegisterRoutes(publicGroup, protectedGroup)

	// Start the Gin server
	port := config.AppConfig.ApiPort
	utils.LogSuccess(fmt.Sprintf("Server started on http://localhost:%s%s\n", port, apiVersion))
	err := router.Run(":" + port)
	if err != nil {
		fmt.Printf("Failed to start the server on port %s: %v\n", port, err)
	}
}
