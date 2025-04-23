package controllers

import (
	"net/http"
	"project1/database"
	"project1/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetCategories получает список категорий с поддержкой пагинации и фильтрации
func GetCategories(c *gin.Context) {
	var pagination models.PaginationParams
	var filter models.CategoryFilter

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

	var categories []models.Category
	query := database.DB.Model(&models.Category{})

	// Применяем фильтры, если они указаны
	if filter.Name != "" {
		query = query.Where("name ILIKE ?", "%"+filter.Name+"%")
	}

	// Подсчитываем общее количество записей
	var total int64
	query.Count(&total)

	// Применяем пагинацию
	offset := (pagination.Page - 1) * pagination.PageSize
	result := query.Offset(offset).Limit(pagination.PageSize).Find(&categories)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch categories"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": categories,
		"meta": gin.H{
			"page":  pagination.Page,
			"limit": pagination.PageSize,
			"total": total,
		},
	})
}

// GetCategoryByID получает информацию о категории по ID
func GetCategoryByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID"})
		return
	}

	var category models.Category
	if err := database.DB.First(&category, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}

	c.JSON(http.StatusOK, category)
}

// CreateCategory создает новую категорию
func CreateCategory(c *gin.Context) {
	var input models.CategoryInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Проверяем, существует ли категория с таким именем
	var existingCategory models.Category
	if result := database.DB.Where("name = ?", input.Name).First(&existingCategory); result.RowsAffected > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Category with this name already exists"})
		return
	}

	category := models.Category{
		Name:        input.Name,
		Description: input.Description,
	}

	if err := database.DB.Create(&category).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create category"})
		return
	}

	c.JSON(http.StatusCreated, category)
}

// UpdateCategory обновляет существующую категорию
func UpdateCategory(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID"})
		return
	}

	var category models.Category
	if err := database.DB.First(&category, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}

	var input models.CategoryInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Проверяем уникальность имени категории, если оно изменилось
	if input.Name != category.Name {
		var existingCategory models.Category
		if result := database.DB.Where("name = ?", input.Name).First(&existingCategory); result.RowsAffected > 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Category with this name already exists"})
			return
		}
	}

	category.Name = input.Name
	category.Description = input.Description

	if err := database.DB.Save(&category).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update category"})
		return
	}

	c.JSON(http.StatusOK, category)
}

// DeleteCategory удаляет категорию
func DeleteCategory(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID"})
		return
	}

	var category models.Category
	if err := database.DB.First(&category, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}

	if err := database.DB.Delete(&category).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete category"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Category deleted successfully"})
}
