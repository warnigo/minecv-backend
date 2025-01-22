package localization

import (
	"path/filepath"

	"github.com/BurntSushi/toml"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

// I18n struct for localization
type I18n struct {
	bundle    *i18n.Bundle
	localizer *i18n.Localizer
}

// NewI18n initializes a new I18n instance
func NewI18n(defaultLang string) *I18n {
	// Create a new bundle with the default language
	bundle := i18n.NewBundle(language.English)

	// Register TOML as the file format
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)

	// Load all localization files from the locales directory
	files, err := filepath.Glob("locales/*.toml")
	if err != nil {
		panic("Failed to find localization files: " + err.Error())
	}

	for _, file := range files {
		_, err := bundle.LoadMessageFile(file)
		if err != nil {
			panic("Failed to load message file: " + file + " - " + err.Error())
		}
	}

	// Initialize localizer with the default language
	localizer := i18n.NewLocalizer(bundle, defaultLang)

	return &I18n{
		bundle:    bundle,
		localizer: localizer,
	}
}

// Translate translates a message ID into the desired language
func (i *I18n) Translate(lang string, id string, templateData map[string]interface{}) string {
	localizer := i18n.NewLocalizer(i.bundle, lang)
	translated, err := localizer.Localize(&i18n.LocalizeConfig{
		MessageID:    id,
		TemplateData: templateData,
	})
	if err != nil {
		return "Translation not found: " + id
	}
	return translated
}
