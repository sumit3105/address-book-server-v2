package routes

import (
	"address-book-server-v2/internal/controllers"
	"address-book-server-v2/internal/core/config"
	"address-book-server-v2/internal/core/middlewares"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)


func setUpAddressRoutes(r *gin.Engine, serverCfg *config.ServerConfig, smtpCfg *config.SMTPConfig, db *gorm.DB) *gin.Engine {

	addressController := controllers.NewAddressController(serverCfg, smtpCfg, db)

	addresses := r.Group("/addresses")
	addresses.Use(middlewares.AuthMiddleware(serverCfg))
	// addresses.Use(middlewares.EnsureUserExistsMiddleware(userRepo))
	{
		addresses.POST("", addressController.Create)
		addresses.GET("", addressController.GetAll)
		addresses.GET("/filter", addressController.GetFiltered)
		addresses.GET("/:id", addressController.GetByID)
		addresses.POST("/export/", addressController.ExportCustom)
		addresses.PUT("/:id", addressController.Update)
		addresses.DELETE("/:id", addressController.Delete)
	}

	return r
}