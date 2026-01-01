package routes

import (
	"address-book-server-v2/internal/controllers"
	"address-book-server-v2/internal/core/config"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func setupAuthRoutes(r *gin.Engine, serverCfg *config.ServerConfig, db *gorm.DB) *gin.Engine {

	authController := controllers.NewAuthController(serverCfg, db)

	auth := r.Group("/auth")
	{
		auth.POST("/register", authController.Register)
		auth.POST("/login", authController.Login)
	}

	return r
}