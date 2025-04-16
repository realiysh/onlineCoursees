package main

import (
	"project1/database"
	"project1/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Инициализируем базу данных
	database.ConnectDB()

	// Создаём сервер Gin
	r := gin.Default()

	// Регистрируем маршруты
	routes.RegisterRoutes(r)

	// Запускаем сервер
	r.Run(":8080")
}
