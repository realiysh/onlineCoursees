package main

import (
	"course-service/database"
	"course-service/middleware"
	"course-service/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// ✅ Шаг 1. Подключаемся к базе данных
	if err := database.Connect(); err != nil {
		log.Fatalf("Ошибка при подключении к БД: %v", err)
	}

	// ✅ Шаг 2. Инициализируем Gin и middleware
	r := gin.New()
	r.Use(middleware.LoggingMiddleware())

	// ✅ Шаг 3. Подключаем маршруты
	routes.SetupRoutes(r)

	// ✅ Шаг 4. Запускаем сервер
	r.Run(":8084")
}
