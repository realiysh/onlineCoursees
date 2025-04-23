package controllers

import (
	"net/http"
	"project1/database"
	"project1/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetAuthors получает список авторов с поддержкой пагинации и фильтрации
func GetAuthors(c *gin.Context) {
	var pagination models.PaginationParams
	var filter models.AuthorFilter

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

	var authors []models.Author
	query := database.DB.Model(&models.Author{})

	// Применяем фильтры, если они указаны
	if filter.Name != "" {
		query = query.Where("name ILIKE ?", "%"+filter.Name+"%")
	}

	// Подсчитываем общее количество записей
	var total int64
	query.Count(&total)

	// Применяем пагинацию
	offset := (pagination.Page - 1) * pagination.PageSize
	result := query.Offset(offset).Limit(pagination.PageSize).Find(&authors)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch authors"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": authors,
		"meta": gin.H{
			"page":  pagination.Page,
			"limit": pagination.PageSize,
			"total": total,
		},
	})
}

// GetAuthorByID получает информацию о авторе по ID
func GetAuthorByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid author ID"})
		return
	}

	var author models.Author
	if err := database.DB.First(&author, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Author not found"})
		return
	}

	c.JSON(http.StatusOK, author)
}

// CreateAuthor создает нового автора
func CreateAuthor(c *gin.Context) {
	var input models.AuthorInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	author := models.Author{
		Name: input.Name,
	}

	if err := database.DB.Create(&author).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create author"})
		return
	}

	c.JSON(http.StatusCreated, author)
}

// UpdateAuthor обновляет существующего автора
func UpdateAuthor(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid author ID"})
		return
	}

	var author models.Author
	if err := database.DB.First(&author, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Author not found"})
		return
	}

	var input models.AuthorInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	author.Name = input.Name

	if err := database.DB.Save(&author).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update author"})
		return
	}

	c.JSON(http.StatusOK, author)
}

// DeleteAuthor удаляет автора
func DeleteAuthor(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid author ID"})
		return
	}

	var author models.Author
	if err := database.DB.First(&author, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Author not found"})
		return
	}

	// Проверяем, используется ли автор в курсах
	var count int64
	if err := database.DB.Model(&models.Course{}).Where("author_id = ?", id).Count(&count).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check related courses"})
		return
	}

	if count > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cannot delete author with related courses"})
		return
	}

	if err := database.DB.Delete(&author).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete author"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Author deleted successfully"})
}
