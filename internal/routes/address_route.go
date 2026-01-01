package routes

import (
	"address-book-server-v2/internal/core/config"
	"address-book-server-v2/internal/core/middlewares"

	"github.com/gin-gonic/gin"
)


func setUpAddressRoutes(r *gin.Engine, serverCfg *config.ServerConfig) *gin.Engine {
	// addressRepo := repositories.NewAddressRepository(db.DB)
	// addressService := services.NewAddressService(addressRepo)
	// addressController := controllers.NewAddressController(addressService, serverCfg)

	addresses := r.Group("/addresses")
	addresses.Use(middlewares.AuthMiddleware(serverCfg))
	// addresses.Use(middlewares.EnsureUserExistsMiddleware(userRepo))
	{
		// addresses.POST("", addressController.Create)
		// addresses.GET("", addressController.GetAll)
		// addresses.GET("/filter", addressController.GetFiltered)
		// addresses.GET("/:id", addressController.GetByID)
		// addresses.GET("/export/custom", addressController.ExportCustom)
		// addresses.GET("/export/sync", addressController.Export)
		// addresses.GET("/export/async", addressController.ExportAsync)
		// addresses.PUT("/:id", addressController.Update)
		// addresses.DELETE("/:id", addressController.Delete)
	}

	return r
}