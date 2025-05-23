package controllers

import (
	"course-service/database"
	"course-service/models"
	"course-service/resty"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetCourses(c *gin.Context) {
	var courses []models.Course
	if err := database.DB.Preload("Category").Find(&courses).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve courses"})
		return
	}
	c.JSON(http.StatusOK, courses)
}

func CreateCourse(c *gin.Context) {
	var input models.CourseInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	course := models.Course{
		Title:       input.Title,
		Description: input.Description,
		Price:       input.Price,
		CategoryID:  input.CategoryID,
	}

	if err := database.DB.Create(&course).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create course"})
		return
	}

	c.JSON(http.StatusCreated, course)
}

func GetCourseByID(c *gin.Context) {
	id := c.Param("id")
	var course models.Course
	if err := database.DB.Preload("Category").First(&course, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Course not found"})
		return
	}
	c.JSON(http.StatusOK, course)
}

func DeleteCourse(c *gin.Context) {
	id := c.Param("id")
	if err := database.DB.Delete(&models.Course{}, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Course not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Course deleted"})
}
func ExampleGetUser(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "нет токена"})
		return
	}

	user, err := resty.Useruser(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "не удалось получить пользователя"})
		return
	}

	c.JSON(http.StatusOK, user)
}
