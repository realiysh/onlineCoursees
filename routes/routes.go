package routes

import (
	"user-service/controllers"
	"user-service/database"
	"user-service/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	authController := controllers.NewAuthController(database.DB)

	api := r.Group("/api")

	// Публичные маршруты
	api.POST("/auth/register", authController.Register)
	api.POST("/auth/login", authController.Login)

	// Защищенные маршруты
	protected := api.Group("")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/users/profile", controllers.GetProfile)
		protected.PUT("/users/profile", controllers.UpdateProfile)
		protected.DELETE("/users/profile", controllers.DeleteProfile)
	}
}
