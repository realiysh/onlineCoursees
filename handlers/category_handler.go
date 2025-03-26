package handlers

import (
	"WelcomeGo/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

var categories = []models.Category{}

func GetCategories(c *gin.Context) {
	c.JSON(http.StatusOK, categories)
}

func AddCategory(c *gin.Context) {
	var newCategory models.Category
	if err := c.ShouldBindJSON(&newCategory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newCategory.ID = len(categories) + 1
	categories = append(categories, newCategory)
	c.JSON(http.StatusCreated, newCategory)
}
