package routes

import (
	"course-service/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		// Курсы
		api.GET("/courses", controllers.GetCourses)
		api.POST("/courses", controllers.CreateCourse)
		api.GET("/courses/:id", controllers.GetCourseByID)
		api.DELETE("/courses/:id", controllers.DeleteCourse)

		// Категории
		api.GET("/categories", controllers.GetCategories)
		api.POST("/categories", controllers.CreateCategory)
		api.GET("/categories/:id", controllers.GetCategoryByID)
		api.PUT("/categories/:id", controllers.UpdateCategory)
		api.DELETE("/categories/:id", controllers.DeleteCategory)

		// Статистика
		api.GET("/stats/courses", controllers.GetCourseStats)
		api.GET("/stats/categories", controllers.GetCategoryStats)
		api.GET("/stats/price-ranges", controllers.GetPriceRangeStats)

		// Поиск
		api.GET("/search/courses", controllers.SearchCourses)
		api.GET("/search/price-range", controllers.SearchByPriceRange)
		api.GET("/search/category/:category_id", controllers.SearchByCategory)
		api.GET("/search/popular", controllers.SearchPopularCourses)

		api.GET("/whoami", controllers.ExampleGetUser)
	}
}
