package controllers

import (
	"course-service/database"
	"course-service/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// SearchCourses выполняет поиск курсов
func SearchCourses(c *gin.Context) {
	query := c.Query("q")
	if query == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Query parameter 'q' is required"})
		return
	}

	var pagination models.PaginationParams
	pagination.Page = 1
	pagination.PageSize = 10

	if err := c.ShouldBindQuery(&pagination); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var courses []models.Course
	searchQuery := database.DB.Model(&models.Course{}).Where("title ILIKE ?", "%"+query+"%")

	var total int64
	searchQuery.Count(&total)

	offset := (pagination.Page - 1) * pagination.PageSize
	if err := searchQuery.Offset(offset).Limit(pagination.PageSize).Find(&courses).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to search courses"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": courses,
		"meta": gin.H{
			"page":      pagination.Page,
			"page_size": pagination.PageSize,
			"total":     total,
			"query":     query,
		},
	})
}

// SearchAuthors выполняет поиск авторов
func SearchAuthors(c *gin.Context) {
	query := c.Query("q")
	if query == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Query parameter 'q' is required"})
		return
	}

	var pagination models.PaginationParams
	pagination.Page = 1
	pagination.PageSize = 10

	if err := c.ShouldBindQuery(&pagination); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var authors []models.Author
	searchQuery := database.DB.Model(&models.Author{}).Where("name ILIKE ?", "%"+query+"%")

	var total int64
	searchQuery.Count(&total)

	offset := (pagination.Page - 1) * pagination.PageSize
	if err := searchQuery.Offset(offset).Limit(pagination.PageSize).Find(&authors).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to search authors"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": authors,
		"meta": gin.H{
			"page":      pagination.Page,
			"page_size": pagination.PageSize,
			"total":     total,
			"query":     query,
		},
	})
}
