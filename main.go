package main

import (
	"log"
	"user-service/database"
	"user-service/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	// Подключение к базе данных
	if err := database.Connect(); err != nil {
		log.Fatalf("Ошибка подключения к БД: %v", err)
	}

	// Инициализация Gin
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.LoggingMiddleware())

	// Роуты
	setupRoutes(r)

	// Запуск сервера
	if err := r.Run(":8084"); err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}
