package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"minecv/internal/infrastructure/localization"
	"minecv/pkg/lib"
	"minecv/pkg/utils"
)

// AuthMiddleware validates the JWT token from the Authorization header and ensures user authentication.
func AuthMiddleware(i18n *localization.I18n) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get user's language preference
		lang := c.GetString("lang")
		translate := utils.GetTranslator(i18n, lang)

		// Get Authorization Header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			utils.RespondError(c, http.StatusUnauthorized, translate("error.unauthorized", nil))
			c.Abort()
			return
		}

		// Extract token
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// Validate JWT
		claims, err := lib.ParseToken(tokenString)
		if err != nil {
			utils.RespondError(c, http.StatusUnauthorized, translate("error.invalid_token", nil))
			c.Abort()
			return
		}

		// Store user ID in context
		c.Set("user_id", claims.UserId)

		c.Next()
	}
}
