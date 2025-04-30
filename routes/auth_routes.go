package routes

import (
	"user-service/controllers"
	"user-service/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupAuthRoutes(r *gin.Engine, db *gorm.DB) {
	authController := controllers.NewAuthController(db)

	auth := r.Group("/auth")
	{
		auth.POST("/register", authController.Register)
		auth.POST("/login", authController.Login)
	}

	protected := r.Group("/api")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/profile", func(c *gin.Context) {
			userID := c.GetUint("user_id")
			email := c.GetString("email")
			role := c.GetString("role")

			c.JSON(200, gin.H{
				"user_id": userID,
				"email":   email,
				"role":    role,
			})
		})
	}
}
