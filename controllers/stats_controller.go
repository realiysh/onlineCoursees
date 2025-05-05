package controllers

import (
	"course-service/database"
	"course-service/models"
	"net/http"

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
	var totalCategories int64
	var avgPrice float64

	database.DB.Model(&models.Course{}).Count(&totalCourses)
	database.DB.Model(&models.Category{}).Count(&totalCategories)
	database.DB.Model(&models.Course{}).Select("AVG(price)").Scan(&avgPrice)

	c.JSON(http.StatusOK, gin.H{
		"total_courses":    totalCourses,
		"total_categories": totalCategories,
		"average_price":    avgPrice,
	})
}

func GetCategoryStats(c *gin.Context) {
	var stats []struct {
		CategoryName string
		CourseCount  int64
		AvgPrice     float64
	}

	database.DB.Model(&models.Category{}).
		Select("categories.name as category_name, COUNT(courses.id) as course_count, AVG(courses.price) as avg_price").
		Joins("LEFT JOIN courses ON courses.category_id = categories.id").
		Group("categories.id, categories.name").
		Scan(&stats)

	c.JSON(http.StatusOK, stats)
}

func GetPriceRangeStats(c *gin.Context) {
	var stats struct {
		MinPrice    float64
		MaxPrice    float64
		PriceRanges []struct {
			Range string
			Count int64
		}
	}

	database.DB.Model(&models.Course{}).
		Select("MIN(price) as min_price, MAX(price) as max_price").
		Scan(&stats)

	ranges := []struct {
		min   float64
		max   float64
		label string
	}{
		{0, 50, "0-50"},
		{50, 100, "50-100"},
		{100, 200, "100-200"},
		{200, 500, "200-500"},
		{500, 1000, "500-1000"},
		{1000, 999999, "1000+"},
	}

	for _, r := range ranges {
		var count int64
		database.DB.Model(&models.Course{}).
			Where("price >= ? AND price < ?", r.min, r.max).
			Count(&count)

		stats.PriceRanges = append(stats.PriceRanges, struct {
			Range string
			Count int64
		}{
			Range: r.label,
			Count: count,
		})
	}

	c.JSON(http.StatusOK, stats)
}
