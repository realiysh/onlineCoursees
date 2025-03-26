package handlers

import (
	"WelcomeGo/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

var authors = []models.Author{}

func GetAuthors(c *gin.Context) {
	c.JSON(http.StatusOK, authors)
}

func AddAuthor(c *gin.Context) {
	var newAuthor models.Author
	if err := c.ShouldBindJSON(&newAuthor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newAuthor.ID = len(authors) + 1
	authors = append(authors, newAuthor)
	c.JSON(http.StatusCreated, newAuthor)
}
