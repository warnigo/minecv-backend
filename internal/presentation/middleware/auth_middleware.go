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
		// Get the user's preferred language
		lang := c.GetString("lang")

		// Translate missing or invalid token message
		tranlate := utils.GetTranslator(i18n, lang)

		// Get the token from the Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			utils.RespondError(c, http.StatusUnauthorized, tranlate("error.unauthorized", nil))
			c.Abort()
			return
		}

		// Remove the "Bearer " prefix
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// Validate the JWT
		claims, err := lib.ParseToken(tokenString)
		if err != nil {
			utils.RespondError(c, http.StatusUnauthorized, tranlate("error.invalid_token", nil))
			c.Abort()
			return
		}

		// Store the user ID in the context for future use
		c.Set("user_id", claims.UserId)

		c.Next()
	}
}
