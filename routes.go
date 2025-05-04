package routes

import (
	"course-service/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.GET("/courses", controllers.GetCourses)
		api.POST("/courses", controllers.CreateCourse)
		api.GET("/courses/:id", controllers.GetCourseByID)
		api.DELETE("/courses/:id", controllers.DeleteCourse)
	}
}
