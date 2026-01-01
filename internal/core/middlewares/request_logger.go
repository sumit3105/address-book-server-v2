package middlewares

import (
	"address-book-server-v2/internal/common/log"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func ReuqestLogger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()

		ctx.Next()

		latency := time.Since(start)		

		logger.Logger.Info(
			"HTTP Request",
			zap.String("method", ctx.Request.Method),
			zap.String("path", ctx.Request.URL.Path),
			zap.Int("status", ctx.Writer.Status()),
			zap.Duration("latency", latency),
			zap.String("ip-address", ctx.ClientIP()),
		)
	}
}