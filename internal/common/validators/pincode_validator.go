package validators

import (
    "regexp"

    "github.com/gin-gonic/gin/binding"
    "github.com/go-playground/validator/v10"
)

func registerPincodeValidator() {
    if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
        v.RegisterValidation("pincode", func(fl validator.FieldLevel) bool {
            pincode := fl.Field().String()

            // Indian PIN code: 6 digits, first digit 1–9
            regex := regexp.MustCompile(`^[1-9][0-9]{5}$`)
            return regex.MatchString(pincode)
        })
    }
}
