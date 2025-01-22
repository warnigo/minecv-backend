package middleware

import (
	"github.com/gin-gonic/gin"

	"minecv/internal/infrastructure/localization"
)

func LocalizationMiddleware(i18n *localization.I18n) gin.HandlerFunc {
	return func(c *gin.Context) {
		lang := c.GetHeader("Accept-Language")
		if lang == "" {
			lang = "en"
		}
		c.Set("lang", lang)
		c.Set("localizer", i18n)
		c.Next()
	}
}