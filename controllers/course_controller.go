package controllers

import (
	"course-service/database"
	"course-service/models"
	"course-service/resty"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CreateCourseRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
}

func GetCourses(c *gin.Context) {
	var courses []models.Course
	if err := database.DB.Find(&courses).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch courses"})
		return
	}
	c.JSON(http.StatusOK, courses)
}

func GetCourseByID(c *gin.Context) {
	id := c.Param("id")
	var course models.Course
	if err := database.DB.First(&course, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Course not found"})
		return
	}
	c.JSON(http.StatusOK, course)
}

func CreateCourse(c *gin.Context) {
	var req CreateCourseRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userIDStr := c.Query("userID")
	userID, err := strconv.ParseUint(userIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// Проверяем существование пользователя
	user, err := resty.GetUserByID(uint(userID))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user"})
		return
	}

	course := &models.Course{
		Title:       req.Title,
		Description: req.Description,
		UserID:      uint(userID),
	}

	if err := database.DB.Create(course).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create course"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Course created",
		"course":  course,
		"user":    user,
	})
}

func UpdateCourse(c *gin.Context) {
	id := c.Param("id")
	var course models.Course
	if err := database.DB.First(&course, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Course not found"})
		return
	}

	var req CreateCourseRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	course.Title = req.Title
	course.Description = req.Description

	if err := database.DB.Save(&course).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update course"})
		return
	}

	c.JSON(http.StatusOK, course)
}

func DeleteCourse(c *gin.Context) {
	id := c.Param("id")
	var course models.Course
	if err := database.DB.First(&course, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Course not found"})
		return
	}

	if err := database.DB.Delete(&course).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete course"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Course deleted successfully"})
}
