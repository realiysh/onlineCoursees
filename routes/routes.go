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

		// Публичные маршруты для просмотра (для упрощения тестирования)
		api.GET("/courses", controllers.GetCourses)
		api.GET("/courses/:id", controllers.GetCourseByID)
		api.GET("/authors", controllers.GetAuthors)
		api.GET("/authors/:id", controllers.GetAuthorByID)
		api.GET("/categories", controllers.GetCategories)
		api.GET("/categories/:id", controllers.GetCategoryByID)
		api.GET("/search/courses", controllers.SearchCourses)
		api.GET("/search/authors", controllers.SearchAuthors)

		// Защищенные маршруты
		protected := api.Group("")
		protected.Use(middleware.AuthMiddleware())
		{
			// Курсы
			protected.POST("/courses", controllers.CreateCourse)
			protected.PUT("/courses/:id", controllers.UpdateCourse)
			protected.DELETE("/courses/:id", controllers.DeleteCourse)

			// Авторы
			protected.POST("/authors", controllers.CreateAuthor)
			protected.PUT("/authors/:id", controllers.UpdateAuthor)
			protected.DELETE("/authors/:id", controllers.DeleteAuthor)

			// Пользователи
			protected.GET("/users", controllers.GetUsers)
			protected.GET("/users/:id", controllers.GetUserByID)
			protected.GET("/users/profile", controllers.GetUserProfile)
			protected.PUT("/users/profile", controllers.UpdateUserProfile)

			// Статистика
			protected.GET("/stats/users", controllers.GetUserStats)
			protected.GET("/stats/courses", controllers.GetCourseStats)

			// Категории
			protected.POST("/categories", controllers.CreateCategory)
			protected.PUT("/categories/:id", controllers.UpdateCategory)
			protected.DELETE("/categories/:id", controllers.DeleteCategory)
		}
	}
}
