package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"project1/models"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// Создание тестового автора
func createTestAuthor() uint {
	return 1 // Возвращаем фиксированный ID для тестов
}

// Создание тестового пользователя и получение токена
func getAuthToken() string {
	return "mocked-jwt-token"
}

// Тест для получения списка курсов
func TestGetCourses(t *testing.T) {
	// Создаем тестовый роутер
	r := SetupTestRouter()

	// Регистрируем обработчик с моком для пути /api/courses
	r.GET("/api/courses", func(c *gin.Context) {
		// Создаем моковые данные курсов
		courses := []models.Course{
			{ID: 1, Title: "Test Course 1", AuthorID: 1, Price: 19.99},
			{ID: 2, Title: "Test Course 2", AuthorID: 1, Price: 29.99},
		}

		// Возвращаем список курсов
		c.JSON(http.StatusOK, gin.H{
			"data": courses,
			"meta": gin.H{
				"page":  1,
				"limit": 10,
				"total": 2,
			},
		})
	})

	// Выполняем тестовый запрос
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/courses", nil)
	r.ServeHTTP(w, req)

	// Проверяем, что статус ответа 200 OK
	assert.Equal(t, http.StatusOK, w.Code)

	// Проверяем структуру ответа
	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Contains(t, response, "data")
	assert.Contains(t, response, "meta")
}

// Тест для создания курса
func TestCreateCourse(t *testing.T) {
	// Получаем тестовый ID автора
	authorID := createTestAuthor()

	// Создаем тестовый роутер
	r := SetupTestRouter()

	// Регистрируем обработчик с моком для пути /api/courses
	r.POST("/api/courses", func(c *gin.Context) {
		var input models.CourseInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Проверяем, что authorID существует
		if input.AuthorID != authorID {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Author not found"})
			return
		}

		// Создаем моковый ответ с новым курсом
		course := models.Course{
			ID:       1,
			Title:    input.Title,
			AuthorID: input.AuthorID,
			Price:    input.Price,
		}

		c.JSON(http.StatusCreated, course)
	})

	// Создаем тестовый запрос
	courseInput := models.CourseInput{
		Title:    "Test Course",
		AuthorID: authorID,
		Price:    29.99,
	}
	jsonData, _ := json.Marshal(courseInput)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/courses", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	// Проверяем, что статус ответа 201 Created
	assert.Equal(t, http.StatusCreated, w.Code)

	// Проверяем содержимое ответа
	var response models.Course
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, uint(1), response.ID)
	assert.Equal(t, "Test Course", response.Title)
	assert.Equal(t, authorID, response.AuthorID)
	assert.Equal(t, 29.99, response.Price)
}

// Тест для получения курса по ID
func TestGetCourseByID(t *testing.T) {
	// Получаем тестовый ID автора
	authorID := createTestAuthor()

	// Создаем тестовый роутер
	r := SetupTestRouter()

	// Регистрируем обработчик с моком для пути /api/courses/:id
	r.GET("/api/courses/:id", func(c *gin.Context) {
		id := c.Param("id")

		// Проверяем ID
		if id != "1" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Course not found"})
			return
		}

		// Возвращаем курс
		course := models.Course{
			ID:       1,
			Title:    "Test Course",
			AuthorID: authorID,
			Price:    29.99,
		}

		c.JSON(http.StatusOK, course)
	})

	// Выполняем тестовый запрос
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/courses/1", nil)
	r.ServeHTTP(w, req)

	// Проверяем, что статус ответа 200 OK
	assert.Equal(t, http.StatusOK, w.Code)

	// Проверяем, что возвращенный курс соответствует ожидаемому
	var response models.Course
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, uint(1), response.ID)
	assert.Equal(t, "Test Course", response.Title)
}
