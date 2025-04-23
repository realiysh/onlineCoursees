package main

import (
	"log"
	"project1/database"
	"project1/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Инициализируем базу данных и запускаем миграции
	log.Println("Initializing database connection...")
	err := database.ConnectDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
		return
	}
	log.Println("Database connected successfully")

	// Создаём сервер Gin
	r := gin.Default()

	// Регистрируем маршруты
	routes.RegisterRoutes(r)

	// Запускаем сервер
	log.Println("Starting server on port 8080...")
	r.Run(":8080")
}
