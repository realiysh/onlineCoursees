package controllers

import (
	"net/http"
	"user-service/database"
	"user-service/models"

	"github.com/gin-gonic/gin"
)

// GetAuthors возвращает список всех авторов
func GetAuthors(c *gin.Context) {
	var authors []models.Author
	if err := database.DB.Find(&authors).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при получении списка авторов"})
		return
	}
	c.JSON(http.StatusOK, authors)
}

// GetAuthorByID возвращает информацию об авторе по ID
func GetAuthorByID(c *gin.Context) {
	id := c.Param("id")
	var author models.Author
	if err := database.DB.First(&author, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Автор не найден"})
		return
	}
	c.JSON(http.StatusOK, author)
}

// CreateAuthor создает нового автора
func CreateAuthor(c *gin.Context) {
	var author models.Author
	if err := c.ShouldBindJSON(&author); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Create(&author).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при создании автора"})
		return
	}

	c.JSON(http.StatusCreated, author)
}

// UpdateAuthor обновляет информацию об авторе
func UpdateAuthor(c *gin.Context) {
	id := c.Param("id")
	var author models.Author
	if err := database.DB.First(&author, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Автор не найден"})
		return
	}

	if err := c.ShouldBindJSON(&author); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Save(&author).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при обновлении автора"})
		return
	}

	c.JSON(http.StatusOK, author)
}

// DeleteAuthor удаляет автора
func DeleteAuthor(c *gin.Context) {
	id := c.Param("id")
	if err := database.DB.Delete(&models.Author{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при удалении автора"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Автор успешно удален"})
}
