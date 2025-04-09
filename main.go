package main

import (
	"github.com/gin-gonic/gin"
	"project1/database"
	"project1/routes"
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
