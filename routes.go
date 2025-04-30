package routes

import (
	"github.com/gin-gonic/gin"
	"project1/controllers"
)

func RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{

		api.POST("/register", controllers.Register)
		api.POST("/login", controllers.Login)

		api.GET("/courses", controllers.GetCourses)
		api.POST("/courses", controllers.CreateCourse)
		api.GET("/courses/:id", controllers.GetCourseByID)
		api.DELETE("/courses/:id", controllers.DeleteCourse)

		api.GET("/authors", controllers.GetAuthors)
		api.POST("/authors", controllers.CreateAuthor)
	}
}
