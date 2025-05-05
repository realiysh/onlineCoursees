package main

import (
	"course-service/database"
	"course-service/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	if err := database.RunMigrations(); err != nil {
		panic("Failed to run migrations: " + err.Error())
	}

	database.ConnectDB()

	r := gin.Default()

	routes.RegisterRoutes(r)

	r.Run(":8080")
}
