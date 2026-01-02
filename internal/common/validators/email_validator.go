package validators

import (
	logger "address-book-server-v2/internal/common/log"
	"fmt"
	"regexp"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

func registerEmailValidator() {
	engine := binding.Validator.Engine()
	logger.Logger.Info(
		"Validator engine type",
		zap.String("type", fmt.Sprintf("%T", engine)),
	)

	if v, ok := engine.(*validator.Validate); ok {
		v.RegisterValidation("strict_email", func(fl validator.FieldLevel) bool {
			email := fl.Field().String()
			regex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
			return regex.MatchString(email)
		})
	} else {
		logger.Logger.Error(
			"Validator engine cast failed",
		)
	}
}
