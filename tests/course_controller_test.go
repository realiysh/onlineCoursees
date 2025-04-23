package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"project1/controllers"
	"project1/database"
	"project1/middleware"
	"project1/models"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

// Создание тестового роутера для тестов
func SetupTestRouter() *gin.Engine {
	// Устанавливаем режим тестирования для Gin
	gin.SetMode(gin.TestMode)

	// Создаем новый роутер
	r := gin.Default()

	return r
}

// Создание тестового автора
func createTestAuthor() uint {
	author := models.Author{
		Name: "Test Author",
	}
	database.DB.Create(&author)
	return author.ID
}

// Создание тестового пользователя и получение токена
func getAuthToken() string {
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

// Тест для получения списка курсов
func TestGetCourses(t *testing.T) {
	// Настраиваем тестовый роутер
	r := SetupTestRouter()
	r.GET("/api/courses", controllers.GetCourses)

	// Выполняем тестовый запрос
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/courses", nil)
	r.ServeHTTP(w, req)

	// Проверяем, что статус ответа 200 OK
	assert.Equal(t, http.StatusOK, w.Code)

	// Проверяем структуру ответа
	var response gin.H
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Contains(t, response, "data")
}

// Тест для создания курса
func TestCreateCourse(t *testing.T) {
	// Создаем тестового автора
	authorID := createTestAuthor()
	defer func() {
		database.DB.Where("id = ?", authorID).Delete(&models.Author{})
	}()

	// Настраиваем тестовый роутер с авторизацией
	r := SetupTestRouter()
	r.POST("/api/courses", func(c *gin.Context) {
		// Эмулируем авторизацию, устанавливая userID вручную
		c.Set("userID", uint(1))
		controllers.CreateCourse(c)
	})

	// Создаем тестовый запрос
	courseJSON := fmt.Sprintf(`{
		"title": "Test Course",
		"description": "Test Course Description",
		"author_id": %d,
		"category_id": 1,
		"published": true
	}`, authorID)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/courses", bytes.NewBufferString(courseJSON))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	// Проверяем, что статус ответа 201 Created
	assert.Equal(t, http.StatusCreated, w.Code)

	// Удаляем тестовый курс
	var response struct {
		Data struct {
			ID uint `json:"id"`
		} `json:"data"`
	}
	json.Unmarshal(w.Body.Bytes(), &response)
	database.DB.Where("id = ?", response.Data.ID).Delete(&models.Course{})
}

// Тест для получения курса по ID
func TestGetCourseByID(t *testing.T) {
	// Создаем тестового автора
	authorID := createTestAuthor()
	defer func() {
		database.DB.Where("id = ?", authorID).Delete(&models.Author{})
	}()

	// Создаем тестовый курс
	course := models.Course{
		Title:       "Test Course",
		Description: "Test Course Description",
		AuthorID:    authorID,
		CategoryID:  1,
		Published:   true,
	}
	database.DB.Create(&course)
	defer database.DB.Where("id = ?", course.ID).Delete(&models.Course{})

	// Настраиваем тестовый роутер
	r := SetupTestRouter()
	r.GET("/api/courses/:id", controllers.GetCourseByID)

	// Выполняем тестовый запрос
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", fmt.Sprintf("/api/courses/%d", course.ID), nil)
	r.ServeHTTP(w, req)

	// Проверяем, что статус ответа 200 OK
	assert.Equal(t, http.StatusOK, w.Code)

	// Проверяем, что возвращенный курс соответствует ожидаемому
	var response struct {
		Data models.Course `json:"data"`
	}
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, course.ID, response.Data.ID)
	assert.Equal(t, course.Title, response.Data.Title)
}
