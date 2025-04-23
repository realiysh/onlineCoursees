package tests

import (
	"project1/database"
	"project1/middleware"
	"project1/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// SetupTestRouter создает тестовый роутер для тестов
func SetupTestRouter() *gin.Engine {
	// Устанавливаем режим тестирования для Gin
	gin.SetMode(gin.TestMode)

	// Создаем новый роутер
	r := gin.Default()

	return r
}

// CreateTestAuthor создает тестового автора для тестов
func CreateTestAuthor() uint {
	author := models.Author{
		Name: "Test Author",
	}
	database.DB.Create(&author)
	return author.ID
}

// GetAuthToken создает тестового пользователя и возвращает JWT токен
func GetAuthToken() string {
	// Создаем пользователя
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("testpassword"), bcrypt.DefaultCost)
	user := models.User{
		Username: "testuser",
		Password: string(hashedPassword),
	}
	database.DB.Create(&user)

	// Генерируем токен
	token, _ := middleware.GenerateToken(user.ID)
	return token
}

// CreateTestCourse создает тестовый курс для тестов
func CreateTestCourse(authorID uint) uint {
	course := models.Course{
		Title:    "Test Course",
		AuthorID: authorID,
		Price:    99.99,
	}
	database.DB.Create(&course)
	return course.ID
}
