package main

import (
	"user-service/controllers"
	"user-service/middleware"

	"github.com/gin-gonic/gin"
)

func setupRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		// Публичные маршруты
		api.POST("/register", controllers.Register)
		api.POST("/login", controllers.Login)
		api.GET("/authors", controllers.GetAuthors)
		api.GET("/authors/:id", controllers.GetAuthorByID)

		// Защищенные маршруты
		protected := api.Group("/")
		protected.Use(middleware.AuthMiddleware())
		{
			protected.GET("/profile", controllers.GetProfile)
			protected.PUT("/profile", controllers.UpdateProfile)
			protected.POST("/authors", controllers.CreateAuthor)
			protected.PUT("/authors/:id", controllers.UpdateAuthor)
			protected.DELETE("/authors/:id", controllers.DeleteAuthor)
		}
	}
}
