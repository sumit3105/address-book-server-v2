package validators

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

func RegisterPincodeValidator(fl validator.FieldLevel) bool {
	pincode := fl.Field().String()

	// Indian PIN code: 6 digits, first digit 1–9
	regex := regexp.MustCompile(`^[1-9][0-9]{5}$`)
	return regex.MatchString(pincode)
}
