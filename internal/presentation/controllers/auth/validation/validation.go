package validation

import (
	"errors"

	"github.com/go-playground/validator/v10"
	"minecv/internal/infrastructure/localization"
)

// ParseValidationErrors maps validation errors to localized messages.
func ParseValidationErrors(err error, i18n *localization.I18n, lang string) map[string]string {
	validationErrors := make(map[string]string)
	var errs validator.ValidationErrors
	if errors.As(err, &errs) {
		for _, fieldErr := range errs {
			normalizedField := normalizeFieldName(fieldErr.Field())
			validationErrors[normalizedField] = validationErrorMessage(fieldErr, i18n, lang)
		}
	} else {
		validationErrors["error"] = i18n.Translate("validations.unknown_error", lang, nil)
	}
	return validationErrors
}

// validationErrorMessage maps validator.FieldError to a localized error message.
func validationErrorMessage(fe validator.FieldError, i18n *localization.I18n, lang string) string {
	switch fe.Tag() {
	case "required":
		return fieldSpecificRequiredMessage(fe, i18n, lang)
	case "email":
		return i18n.Translate("validations.invalid_email", lang, nil)
	case "omitempty":
		return i18n.Translate("validations.optional_field", lang, nil)
	default:
		return i18n.Translate("validations.invalid_field", lang, nil)
	}
}

// fieldSpecificRequiredMessage generates localized messages for required fields.
func fieldSpecificRequiredMessage(fe validator.FieldError, i18n *localization.I18n, lang string) string {
	switch fe.Kind().String() {
	case "string":
		if fe.Tag() == "email" {
			return i18n.Translate("validations.email_required", lang, nil)
		}
		return i18n.Translate("validations.field_required", lang, nil)
	default:
		return i18n.Translate("validations.value_required", lang, nil)
	}
}

// normalizeFieldName converts CamelCase field names to snake_case.
func normalizeFieldName(field string) string {
	var result []rune
	runes := []rune(field)
	for i, r := range runes {
		if r >= 'A' && r <= 'Z' {
			if i > 0 {
				result = append(result, '_')
			}
			r += 'a' - 'A'
		}
		result = append(result, r)
	}
	return string(result)
}
