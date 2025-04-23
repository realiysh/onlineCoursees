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

	"github.com/stretchr/testify/assert"
)

// Тест для получения списка авторов
func TestGetAuthors(t *testing.T) {
	// Подготавливаем тестовую БД
	database.ConnectDB()

	// Получаем токен для авторизации
	token := GetAuthToken()

	// Создаем тестовый роутер
	r := SetupTestRouter()
	authorGroup := r.Group("/api")
	authorGroup.Use(middleware.AuthMiddleware())
	{
		authorGroup.GET("/authors", controllers.GetAuthors)
	}

	// Создаем запрос
	req, _ := http.NewRequest("GET", "/api/authors", nil)
	req.Header.Set("Authorization", "Bearer "+token)

	// Выполняем запрос
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Проверяем ответ
	assert.Equal(t, http.StatusOK, w.Code)

	// Проверяем структуру ответа
	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Contains(t, response, "data")
	assert.Contains(t, response, "meta")

	// Очищаем тестовые данные
	database.DB.Where("username = ?", "testuser").Delete(&models.User{})
}

// Тест для создания автора
func TestCreateAuthor(t *testing.T) {
	// Подготавливаем тестовую БД
	database.ConnectDB()

	// Получаем токен для авторизации
	token := GetAuthToken()

	// Создаем тестовый роутер
	r := SetupTestRouter()
	authorGroup := r.Group("/api")
	authorGroup.Use(middleware.AuthMiddleware())
	{
		authorGroup.POST("/authors", controllers.CreateAuthor)
	}

	// Создаем тестовые данные
	authorInput := models.AuthorInput{
		Name: "Test Author Creation",
	}
	jsonData, _ := json.Marshal(authorInput)

	// Создаем запрос
	req, _ := http.NewRequest("POST", "/api/authors", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	// Выполняем запрос
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Проверяем ответ
	assert.Equal(t, http.StatusCreated, w.Code)

	// Получаем ID созданного автора для удаления
	var response models.Author
	json.Unmarshal(w.Body.Bytes(), &response)

	// Очищаем тестовые данные
	database.DB.Delete(&models.Author{}, response.ID)
	database.DB.Where("username = ?", "testuser").Delete(&models.User{})
}

// Тест для получения автора по ID
func TestGetAuthorByID(t *testing.T) {
	// Подготавливаем тестовую БД
	database.ConnectDB()

	// Создаем тестового автора
	author := models.Author{
		Name: "Test Author GetByID",
	}
	database.DB.Create(&author)

	// Получаем токен для авторизации
	token := GetAuthToken()

	// Создаем тестовый роутер
	r := SetupTestRouter()
	authorGroup := r.Group("/api")
	authorGroup.Use(middleware.AuthMiddleware())
	{
		authorGroup.GET("/authors/:id", controllers.GetAuthorByID)
	}

	// Создаем запрос
	req, _ := http.NewRequest("GET", fmt.Sprintf("/api/authors/%d", author.ID), nil)
	req.Header.Set("Authorization", "Bearer "+token)

	// Выполняем запрос
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Проверяем ответ
	assert.Equal(t, http.StatusOK, w.Code)

	// Проверяем содержимое ответа
	var response models.Author
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, author.ID, response.ID)
	assert.Equal(t, author.Name, response.Name)

	// Очищаем тестовые данные
	database.DB.Delete(&models.Author{}, author.ID)
	database.DB.Where("username = ?", "testuser").Delete(&models.User{})
}

// Тест для обновления автора
func TestUpdateAuthor(t *testing.T) {
	// Подготавливаем тестовую БД
	database.ConnectDB()

	// Создаем тестового автора
	author := models.Author{
		Name: "Test Author Before Update",
	}
	database.DB.Create(&author)

	// Получаем токен для авторизации
	token := GetAuthToken()

	// Создаем тестовый роутер
	r := SetupTestRouter()
	authorGroup := r.Group("/api")
	authorGroup.Use(middleware.AuthMiddleware())
	{
		authorGroup.PUT("/authors/:id", controllers.UpdateAuthor)
	}

	// Создаем тестовые данные
	authorInput := models.AuthorInput{
		Name: "Test Author After Update",
	}
	jsonData, _ := json.Marshal(authorInput)

	// Создаем запрос
	req, _ := http.NewRequest("PUT", fmt.Sprintf("/api/authors/%d", author.ID), bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	// Выполняем запрос
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Проверяем ответ
	assert.Equal(t, http.StatusOK, w.Code)

	// Проверяем, что автор был обновлен
	var updatedAuthor models.Author
	database.DB.First(&updatedAuthor, author.ID)
	assert.Equal(t, "Test Author After Update", updatedAuthor.Name)

	// Очищаем тестовые данные
	database.DB.Delete(&models.Author{}, author.ID)
	database.DB.Where("username = ?", "testuser").Delete(&models.User{})
}

// DeleteAuthor удаляет автора
func TestDeleteAuthor(t *testing.T) {
	// Подготавливаем тестовую БД
	database.ConnectDB()

	// Создаем тестового автора
	author := models.Author{
		Name: "Test Author for Delete",
	}
	database.DB.Create(&author)

	// Получаем токен для авторизации
	token := GetAuthToken()

	// Создаем тестовый роутер
	r := SetupTestRouter()
	authorGroup := r.Group("/api")
	authorGroup.Use(middleware.AuthMiddleware())
	{
		authorGroup.DELETE("/authors/:id", controllers.DeleteAuthor)
	}

	// Создаем запрос
	req, _ := http.NewRequest("DELETE", fmt.Sprintf("/api/authors/%d", author.ID), nil)
	req.Header.Set("Authorization", "Bearer "+token)

	// Выполняем запрос
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Проверяем ответ
	assert.Equal(t, http.StatusOK, w.Code)

	// Проверяем, что автор был удален
	var deletedAuthor models.Author
	result := database.DB.First(&deletedAuthor, author.ID)
	assert.Error(t, result.Error) // Должна быть ошибка, т.к. автор удален

	// Очищаем тестовые данные
	database.DB.Where("username = ?", "testuser").Delete(&models.User{})
}

// Тест на попытку удаления автора, у которого есть связанные курсы
func TestDeleteAuthorWithCourses(t *testing.T) {
	// Подготавливаем тестовую БД
	database.ConnectDB()

	// Создаем тестового автора
	author := models.Author{
		Name: "Test Author with Courses",
	}
	database.DB.Create(&author)

	// Создаем тестовый курс для автора
	course := models.Course{
		Title:    "Test Course",
		AuthorID: author.ID,
		Price:    99.99,
	}
	database.DB.Create(&course)

	// Получаем токен для авторизации
	token := GetAuthToken()

	// Создаем тестовый роутер
	r := SetupTestRouter()
	authorGroup := r.Group("/api")
	authorGroup.Use(middleware.AuthMiddleware())
	{
		authorGroup.DELETE("/authors/:id", controllers.DeleteAuthor)
	}

	// Создаем запрос на удаление автора с курсами
	req, _ := http.NewRequest("DELETE", fmt.Sprintf("/api/authors/%d", author.ID), nil)
	req.Header.Set("Authorization", "Bearer "+token)

	// Выполняем запрос
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Проверяем ответ - должен быть ошибочный статус, т.к. нельзя удалить автора с курсами
	assert.Equal(t, http.StatusBadRequest, w.Code)

	// Проверяем, что автор не был удален
	var checkAuthor models.Author
	result := database.DB.First(&checkAuthor, author.ID)
	assert.NoError(t, result.Error) // Автор должен существовать

	// Очищаем тестовые данные
	database.DB.Delete(&models.Course{}, course.ID)
	database.DB.Delete(&models.Author{}, author.ID)
	database.DB.Where("username = ?", "testuser").Delete(&models.User{})
}
