package routes

import (
	"project1/controllers"
	"project1/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		// Публичные маршруты
		api.POST("/register", controllers.Register)
		api.POST("/login", controllers.Login)

		// Защищенные маршруты
		protected := api.Group("")
		protected.Use(middleware.AuthMiddleware())
		{
			// Курсы
			protected.GET("/courses", controllers.GetCourses)
			protected.POST("/courses", controllers.CreateCourse)
			protected.GET("/courses/:id", controllers.GetCourseByID)
			protected.PUT("/courses/:id", controllers.UpdateCourse)
			protected.DELETE("/courses/:id", controllers.DeleteCourse)

			// Авторы
			protected.GET("/authors", controllers.GetAuthors)
			protected.POST("/authors", controllers.CreateAuthor)
			protected.GET("/authors/:id", controllers.GetAuthorByID)
			protected.PUT("/authors/:id", controllers.UpdateAuthor)
			protected.DELETE("/authors/:id", controllers.DeleteAuthor)

			// Пользователи
			protected.GET("/users/profile", controllers.GetUserProfile)
			protected.PUT("/users/profile", controllers.UpdateUserProfile)
		}
	}
}
