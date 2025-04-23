package tests

import (
	"project1/middleware"
	"project1/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// Тестовая база данных для моков
var TestDB *gorm.DB

// SetupTestRouter создает тестовый роутер для тестов
func SetupTestRouter() *gin.Engine {
	// Устанавливаем режим тестирования для Gin
	gin.SetMode(gin.TestMode)

	// Создаем новый роутер
	r := gin.Default()

	return r
}

// MockAuthor создает моковый объект Author без записи в базу
func MockAuthor() models.Author {
	return models.Author{
		ID:   1,
		Name: "Test Author",
	}
}

// MockCourse создает моковый объект Course без записи в базу
func MockCourse(authorID uint) models.Course {
	return models.Course{
		ID:       1,
		Title:    "Test Course",
		AuthorID: authorID,
		Price:    99.99,
	}
}

// MockUser создает моковый объект User без записи в базу
func MockUser() models.User {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("testpassword"), bcrypt.DefaultCost)
	return models.User{
		ID:       1,
		Username: "testuser",
		Password: string(hashedPassword),
	}
}

// MockAuthToken создает тестовый JWT токен без создания пользователя
func MockAuthToken() string {
	// Генерируем токен для моковой записи с ID=1
	token, _ := middleware.GenerateToken(1)
	return token
}

// Старые функции, требующие реальной базы данных (оставлены для обратной совместимости)

// CreateTestAuthor создает тестового автора для тестов
func CreateTestAuthor() uint {
	author := models.Author{
		Name: "Test Author",
	}
	// database.DB.Create(&author) - закомментировано, т.к. требует базу данных
	return author.ID
}

// GetAuthToken создает тестового пользователя и возвращает JWT токен
func GetAuthToken() string {
	// Вместо работы с базой, используем мок-функцию
	return MockAuthToken()
}

// CreateTestCourse создает тестовый курс для тестов
func CreateTestCourse(authorID uint) uint {
	course := models.Course{
		Title:    "Test Course",
		AuthorID: authorID,
		Price:    99.99,
	}
	// database.DB.Create(&course) - закомментировано, т.к. требует базу данных
	return course.ID
}
