package handlers

import (
	"WelcomeGo/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var books = []models.Book{}

func GetBooks(c *gin.Context) {
	c.JSON(http.StatusOK, books)
}

func GetBookByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	for _, book := range books {
		if book.ID == id {
			c.JSON(http.StatusOK, book)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
}

func AddBook(c *gin.Context) {
	var newBook models.Book
	if err := c.ShouldBindJSON(&newBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newBook.ID = len(books) + 1
	books = append(books, newBook)
	c.JSON(http.StatusCreated, newBook)
}

func UpdateBook(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	for i, book := range books {
		if book.ID == id {
			if err := c.ShouldBindJSON(&books[i]); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, books[i])
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
}

func DeleteBook(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	for i, book := range books {
		if book.ID == id {
			books = append(books[:i], books[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Book deleted"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
}
