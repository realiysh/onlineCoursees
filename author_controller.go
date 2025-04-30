package controllers

import (
	"net/http"
	"project1/database"
	"project1/models"

	"github.com/gin-gonic/gin"
)

func GetAuthors(c *gin.Context) {
	var authors []models.Author
	if err := database.DB.Find(&authors).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve authors"})
		return
	}
	c.JSON(http.StatusOK, authors)
}

func CreateAuthor(c *gin.Context) {
	var author models.Author
	if err := c.ShouldBindJSON(&author); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	if err := database.DB.Create(&author).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create author"})
		return
	}
	c.JSON(http.StatusCreated, author)
}
