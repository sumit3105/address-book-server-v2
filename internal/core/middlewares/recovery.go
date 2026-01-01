package middlewares

import (
	"address-book-server-v2/internal/common/log"
	"address-book-server-v2/internal/common/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func GlobalRecovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				logger.Logger.Error("panic recovered", zap.Any("error", err))
				utils.Error(c, http.StatusInternalServerError, "internal server error")
				c.Abort()
			}
		}()
		c.Next()
	}
}