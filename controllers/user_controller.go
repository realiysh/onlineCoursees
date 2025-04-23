package controllers

import (
	"net/http"
	"project1/database"
	"project1/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetUserProfile получает профиль текущего пользователя
func GetUserProfile(c *gin.Context) {
	// Получаем ID пользователя из контекста (установлен middleware)
	userID, _ := c.Get("user_id")

	var user models.User
	if err := database.DB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Не возвращаем пароль
	user.Password = ""

	c.JSON(http.StatusOK, user)
}

// UpdateUserProfile обновляет профиль текущего пользователя
func UpdateUserProfile(c *gin.Context) {
	// Получаем ID пользователя из контекста (установлен middleware)
	userID, _ := c.Get("user_id")

	var user models.User
	if err := database.DB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	var input models.UpdateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Проверка уникальности имени пользователя
	if input.Username != "" && input.Username != user.Username {
		var existingUser models.User
		if result := database.DB.Where("username = ?", input.Username).First(&existingUser); result.RowsAffected > 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Username already taken"})
			return
		}
		user.Username = input.Username
	}

	if err := database.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update profile"})
		return
	}

	// Не возвращаем пароль
	user.Password = ""

	c.JSON(http.StatusOK, user)
}

// GetUsers получает список всех пользователей с поддержкой пагинации и фильтрации
func GetUsers(c *gin.Context) {
	var pagination models.PaginationParams
	var filter models.UserFilter

	// Устанавливаем параметры пагинации по умолчанию
	pagination.Page = 1
	pagination.PageSize = 10

	// Биндим параметры пагинации из запроса
	if err := c.ShouldBindQuery(&pagination); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Биндим параметры фильтрации из запроса
	if err := c.ShouldBindQuery(&filter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var users []models.User
	query := database.DB.Model(&models.User{})

	// Применяем фильтры, если они указаны
	if filter.Username != "" {
		query = query.Where("username ILIKE ?", "%"+filter.Username+"%")
	}

	// Подсчитываем общее количество записей
	var total int64
	query.Count(&total)

	// Применяем пагинацию
	offset := (pagination.Page - 1) * pagination.PageSize
	result := query.Offset(offset).Limit(pagination.PageSize).Find(&users)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		return
	}

	// Не возвращаем пароли
	for i := range users {
		users[i].Password = ""
	}

	c.JSON(http.StatusOK, gin.H{
		"data": users,
		"meta": gin.H{
			"page":  pagination.Page,
			"limit": pagination.PageSize,
			"total": total,
		},
	})
}

// GetUserByID получает информацию о пользователе по ID
func GetUserByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	var user models.User
	if err := database.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Не возвращаем пароль
	user.Password = ""

	c.JSON(http.StatusOK, user)
}
