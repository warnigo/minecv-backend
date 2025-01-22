package utils

import "minecv/internal/infrastructure/localization"

// TranslatorFunc defines a shorthand function for translating messages
type TranslatorFunc func(messageID string, templateData map[string]interface{}) string

// GetTranslator initializes and returns a shorthand Translate function
func GetTranslator(i18n *localization.I18n, lang string) TranslatorFunc {
	return func(messageID string, templateData map[string]interface{}) string {
		return i18n.Translate(lang, messageID, templateData)
	}
}
