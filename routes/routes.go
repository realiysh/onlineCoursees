package routes

import (
	"user-service/controllers"
	"user-service/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api")

	// Публичные маршруты
	api.POST("/register", controllers.Register)
	api.POST("/login", controllers.Login)

	// Защищённые маршруты
	protected := api.Group("/")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/profile", controllers.GetProfile)
		protected.PUT("/profile", controllers.UpdateProfile)
		protected.DELETE("/profile", controllers.DeleteProfile)

		// CRUD пользователей (только для админов)
		protected.GET("/users", controllers.GetAllUsers)
		protected.GET("/users/:id", controllers.GetUserByID)
		protected.PUT("/users/:id", controllers.UpdateUserByID)
		protected.DELETE("/users/:id", controllers.DeleteUserByID)
	}
}
