package validators

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

func RegisterPhoneValidator(fl validator.FieldLevel) bool {
	phone := fl.Field().String()

	// E.164 compatible (10–15 digits, optional +)
	regex := regexp.MustCompile(`^\+?[1-9]\d{9,14}$`)
	return regex.MatchString(phone)
}
