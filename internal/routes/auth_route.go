package routes

import "github.com/gin-gonic/gin"

func setupAuthRoutes(r *gin.Engine) *gin.Engine {
	
	// userRepo := repositories.NewUserRepository(db.DB)
	// authService := services.NewAuthService(userRepo)
	// authController := controllers.NewAuthController(authService)

	// auth := r.Group("/auth")
	// {
	// 	// auth.POST("/register", authController.Register)
	// 	// auth.POST("/login", authController.Login)
	// }

	return r
}