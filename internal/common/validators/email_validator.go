package validators

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

func RegisterEmailValidator(fl validator.FieldLevel) bool {
	email := fl.Field().String()
	regex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return regex.MatchString(email)
}
