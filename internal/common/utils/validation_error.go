package utils

import (
	"github.com/go-playground/validator/v10"
	"regexp"
	"strings"
)

func FormatValidationErrors(err error) map[string]string {
	errors := make(map[string]string)

	validationErrors, ok := err.(validator.ValidationErrors)
	if !ok {
		return errors
	}

	for _, fieldErr := range validationErrors {
		field := toSnakeCase(fieldErr.Field())

		switch fieldErr.Tag() {

		case "required":
			errors[field] = "This field is required"

		case "email":
			errors[field] = "Invalid email format"

		case "min":
			errors[field] = "Must be at least " + fieldErr.Param() + " characters long"

		case "password":
			errors[field] = "Password must be at least 8 characters and include uppercase, lowercase, number, and special character"

		default:
			errors[field] = "Invalid value"
		}
	}

	return errors
}

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

func toSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}