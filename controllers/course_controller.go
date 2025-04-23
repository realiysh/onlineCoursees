package controllers

import (
	"net/http"
	"project1/database"
	"project1/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetCourses получает список курсов с поддержкой пагинации и фильтрации
func GetCourses(c *gin.Context) {
	var pagination models.PaginationParams
	var filter models.CourseFilter

	// Устанавливаем параметры пагинации по умолчанию
	pagination.Page = 1
	pagination.PageSize = 10

	// Биндим параметры пагинации из запроса
	if err := c.ShouldBindQuery(&pagination); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Биндим параметры фильтрации из запроса
	if err := c.ShouldBindQuery(&filter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var courses []models.Course
	query := database.DB.Model(&models.Course{})

	// Применяем фильтры, если они указаны
	if filter.Title != "" {
		query = query.Where("title ILIKE ?", "%"+filter.Title+"%")
	}
	if filter.AuthorID != 0 {
		query = query.Where("author_id = ?", filter.AuthorID)
	}
	if filter.MinPrice > 0 {
		query = query.Where("price >= ?", filter.MinPrice)
	}
	if filter.MaxPrice > 0 {
		query = query.Where("price <= ?", filter.MaxPrice)
	}

	// Подсчитываем общее количество записей
	var total int64
	query.Count(&total)

	// Применяем пагинацию
	offset := (pagination.Page - 1) * pagination.PageSize
	result := query.Offset(offset).Limit(pagination.PageSize).Find(&courses)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch courses"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": courses,
		"meta": gin.H{
			"page":  pagination.Page,
			"limit": pagination.PageSize,
			"total": total,
		},
	})
}

// GetCourseByID получает информацию о курсе по ID
func GetCourseByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid course ID"})
		return
	}

	var course models.Course
	if err := database.DB.First(&course, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Course not found"})
		return
	}

	c.JSON(http.StatusOK, course)
}

// CreateCourse создает новый курс
func CreateCourse(c *gin.Context) {
	var input models.CourseInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Проверяем, существует ли автор
	var author models.Author
	if err := database.DB.First(&author, input.AuthorID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Author not found"})
		return
	}

	course := models.Course{
		Title:    input.Title,
		AuthorID: input.AuthorID,
		Price:    input.Price,
	}

	if err := database.DB.Create(&course).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create course"})
		return
	}

	c.JSON(http.StatusCreated, course)
}

// UpdateCourse обновляет существующий курс
func UpdateCourse(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid course ID"})
		return
	}

	var course models.Course
	if err := database.DB.First(&course, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Course not found"})
		return
	}

	var input models.CourseInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Если автор обновлен, проверяем его существование
	if input.AuthorID != 0 && input.AuthorID != course.AuthorID {
		var author models.Author
		if err := database.DB.First(&author, input.AuthorID).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Author not found"})
			return
		}
		course.AuthorID = input.AuthorID
	}

	course.Title = input.Title
	course.Price = input.Price

	if err := database.DB.Save(&course).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update course"})
		return
	}

	c.JSON(http.StatusOK, course)
}

// DeleteCourse удаляет курс
func DeleteCourse(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid course ID"})
		return
	}

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
