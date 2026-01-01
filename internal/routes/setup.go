package routes

import (
	"address-book-server-v2/internal/core/application"
	"address-book-server-v2/internal/core/middlewares"

	"github.com/gin-gonic/gin"
)

func Setup(app *application.App) *gin.Engine {
	r := gin.New()
	
	r.Use(middlewares.ReuqestLogger(), middlewares.GlobalRecovery())

	r.Static("/downloads", "./exports")

	r = setupAuthRoutes(r, app.Cfg.ServerCfg, app.DB)

	r = setUpAddressRoutes(r, app.Cfg.ServerCfg, app.Cfg.SMTPCfg, app.DB)

	return r
}