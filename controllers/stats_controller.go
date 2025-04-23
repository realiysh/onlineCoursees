package controllers

import (
	"net/http"
	"project1/database"
	"project1/models"

	"github.com/gin-gonic/gin"
)

// GetUserStats получает статистику по пользователям
func GetUserStats(c *gin.Context) {
	var totalUsers int64
	database.DB.Model(&models.User{}).Count(&totalUsers)

	var newUsersToday int64
	database.DB.Model(&models.User{}).Where("created_at >= NOW() - INTERVAL '1 day'").Count(&newUsersToday)

	c.JSON(http.StatusOK, gin.H{
		"total_users":     totalUsers,
		"new_users_today": newUsersToday,
	})
}

// GetCourseStats получает статистику по курсам
func GetCourseStats(c *gin.Context) {
	var totalCourses int64
	database.DB.Model(&models.Course{}).Count(&totalCourses)

	var newCoursesToday int64
	database.DB.Model(&models.Course{}).Where("created_at >= NOW() - INTERVAL '1 day'").Count(&newCoursesToday)

	// Средняя цена курсов
	type AvgPriceResult struct {
		AvgPrice float64
	}
	var avgPriceResult AvgPriceResult
	database.DB.Model(&models.Course{}).Select("COALESCE(AVG(price), 0) as avg_price").Scan(&avgPriceResult)

	// Топ авторов по количеству курсов
	type AuthorWithCoursesCount struct {
		ID           uint   `json:"id"`
		Name         string `json:"name"`
		CoursesCount int64  `json:"courses_count"`
	}
	var topAuthors []AuthorWithCoursesCount
	database.DB.Table("authors").
		Select("authors.id, authors.name, COUNT(courses.id) as courses_count").
		Joins("JOIN courses ON authors.id = courses.author_id").
		Group("authors.id").
		Order("courses_count DESC").
		Limit(5).
		Scan(&topAuthors)

	c.JSON(http.StatusOK, gin.H{
		"total_courses":     totalCourses,
		"new_courses_today": newCoursesToday,
		"avg_price":         avgPriceResult.AvgPrice,
		"top_authors":       topAuthors,
	})
}
