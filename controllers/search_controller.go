package controllers

import (
	"course-service/database"
	"course-service/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// SearchCourses выполняет поиск курсов
func SearchCourses(c *gin.Context) {
	query := c.Query("q")
	var courses []models.Course

	tx := database.DB.Preload("Category")

	if query != "" {
		tx = tx.Where("title ILIKE ? OR description ILIKE ?", "%"+query+"%", "%"+query+"%")
	}

	if err := tx.Find(&courses).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to search courses"})
		return
	}

	c.JSON(http.StatusOK, courses)
}

func SearchByPriceRange(c *gin.Context) {
	minPrice := c.Query("min_price")
	maxPrice := c.Query("max_price")

	var courses []models.Course
	tx := database.DB.Preload("Category")

	if minPrice != "" {
		tx = tx.Where("price >= ?", minPrice)
	}
	if maxPrice != "" {
		tx = tx.Where("price <= ?", maxPrice)
	}

	if err := tx.Find(&courses).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to filter courses by price range"})
		return
	}

	c.JSON(http.StatusOK, courses)
}

func SearchByCategory(c *gin.Context) {
	categoryID := c.Param("category_id")
	var courses []models.Course

	database.DB.Where("category_id = ?", categoryID).Find(&courses)
	c.JSON(http.StatusOK, courses)
}

func SearchPopularCourses(c *gin.Context) {
	var courses []models.Course
	limitStr := c.DefaultQuery("limit", "10")
	limit := 10 // значение по умолчанию
	fmt.Sscanf(limitStr, "%d", &limit)

	database.DB.Order("price DESC").Limit(limit).Find(&courses)
	c.JSON(http.StatusOK, courses)
}
