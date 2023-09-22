package helpers

import (
	"strconv"

	validation "github.com/go-ozzo/ozzo-validation"
)

func ValidateNonEmptyString(str string) error {
	return validation.Validate(str, validation.Required)
}

func SubmoduleIdentifiersToString(submoduleId int64, submoduleAccessHash *int64) (string, *string) {
	var accessHash *string = nil
	if submoduleAccessHash != nil {
		tmp := strconv.Itoa(int(*submoduleAccessHash))
		accessHash = &tmp
	}

	return strconv.Itoa(int(submoduleId)), accessHash
}
