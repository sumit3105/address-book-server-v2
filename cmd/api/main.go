package main

import (
	logger "address-book-server-v2/internal/common/log"
	"address-book-server-v2/internal/core/application"
	"address-book-server-v2/internal/routes"

	"go.uber.org/zap"
)

func main() {
	app := application.NewApp()
	logger.Logger.Info("App loded", zap.Any("App", app))

	// Router setup
	r := routes.Setup(app)

	// Start server
	r.Run(":" + app.Cfg.ServerCfg.ServerPort)
}
