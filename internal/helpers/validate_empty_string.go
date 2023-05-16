package helpers

import validation "github.com/go-ozzo/ozzo-validation"

func ValidateNonEmptyString(str string) error {
	return validation.Validate(str, validation.Required)
}
