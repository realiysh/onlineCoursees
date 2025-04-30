package main

import (
	"log"

	"user-service/database"
	"user-service/middleware"
	"user-service/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	if err := database.Connect(); err != nil {
		log.Fatalf("Ошибка при подключении к БД: %v", err)
	}

	r := gin.New()
	r.Use(middleware.LoggingMiddleware())

	routes.SetupRoutes(r)

	r.Run(":8083")
}
