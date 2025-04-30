package routes

import (
	"course-service/controllers"
	"course-service/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api")

	// === Публичные маршруты ===

	// Курсы
	api.GET("/courses", controllers.GetCourses)
	api.GET("/courses/:id", controllers.GetCourseByID)

	// Категории
	api.GET("/categories", controllers.GetCategories)
	api.GET("/categories/:id", controllers.GetCategoryByID)

	// Поиск
	api.GET("/search/courses", controllers.SearchCourses)
	api.GET("/search/authors", controllers.SearchAuthors)

	// === Защищённые маршруты ===
	protected := api.Group("")
	protected.Use(middleware.AuthMiddleware())
	{
		// Курсы
		protected.POST("/courses", controllers.CreateCourse)
		protected.PUT("/courses/:id", controllers.UpdateCourse)
		protected.DELETE("/courses/:id", controllers.DeleteCourse)

		// Категории
		protected.POST("/categories", controllers.CreateCategory)
		protected.PUT("/categories/:id", controllers.UpdateCategory)
		protected.DELETE("/categories/:id", controllers.DeleteCategory)

		// Статистика
		protected.GET("/stats/courses", controllers.GetCourseStats)
	}
}
