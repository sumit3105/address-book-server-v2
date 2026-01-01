package main

import (
	logger "address-book-server-v2/internal/common/log"
	"address-book-server-v2/internal/core/application"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	app := application.NewApp()
	logger.Logger.Info("App loded", zap.Any("App", app))

	r := gin.New()
	// Start server
	r.Run(":" + app.Cfg.ServerCfg.ServerPort)
}
