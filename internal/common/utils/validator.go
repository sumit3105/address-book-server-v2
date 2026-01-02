package utils

import (
	"address-book-server-v2/internal/common/validators"

	"github.com/go-playground/validator/v10"
)

var Validate *validator.Validate

func InitValidator() {
	Validate = validator.New()

	// Register custom validators
	Validate.RegisterValidation("strict_email", validators.RegisterEmailValidator)
	Validate.RegisterValidation("phone", validators.RegisterPhoneValidator)
	Validate.RegisterValidation("pincode", validators.RegisterPincodeValidator)
}
