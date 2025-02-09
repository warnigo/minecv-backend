package middleware

import (
	"github.com/gin-gonic/gin"
	"minecv/pkg/utils"

	"minecv/internal/infrastructure/localization"
)

func LocalizationMiddleware(i18n *localization.I18n) gin.HandlerFunc {
	return func(c *gin.Context) {
		lang := c.GetHeader("Accept-Language")
		if lang == "" {
			lang = "en"
		}

		c.Set("translator", utils.GetTranslator(i18n, lang))
		c.Set("i18n", i18n)
		c.Set("lang", lang)

		c.Next()
	}
}
