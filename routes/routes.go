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
			protected.GET("/courses/:id", controllers.GetCourseByID)
			protected.POST("/courses", controllers.CreateCourse)
			protected.PUT("/courses/:id", controllers.UpdateCourse)
			protected.DELETE("/courses/:id", controllers.DeleteCourse)

			// Авторы
			protected.GET("/authors", controllers.GetAuthors)
			protected.GET("/authors/:id", controllers.GetAuthorByID)
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

			// Поиск
			protected.GET("/search/courses", controllers.SearchCourses)
			protected.GET("/search/authors", controllers.SearchAuthors)

			// Категории
			protected.GET("/categories", controllers.GetCategories)
			protected.GET("/categories/:id", controllers.GetCategoryByID)
			protected.POST("/categories", controllers.CreateCategory)
			protected.PUT("/categories/:id", controllers.UpdateCategory)
			protected.DELETE("/categories/:id", controllers.DeleteCategory)
		}
	}
}
