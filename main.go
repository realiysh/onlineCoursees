package main

import (
	"WelcomeGo/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Кітап маршруты
	r.GET("/books", handlers.GetBooks)
	r.GET("/books/:id", handlers.GetBookByID)
	r.POST("/books", handlers.AddBook)
	r.PUT("/books/:id", handlers.UpdateBook)
	r.DELETE("/books/:id", handlers.DeleteBook)

	// Автор маршруты
	r.GET("/authors", handlers.GetAuthors)
	r.POST("/authors", handlers.AddAuthor)

	// Категория маршруты
	r.GET("/categories", handlers.GetCategories)
	r.POST("/categories", handlers.AddCategory)

	r.Run(":8080")
}
