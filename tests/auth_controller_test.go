package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"project1/controllers"
	"project1/database"
	"project1/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Тест для регистрации пользователя
func TestRegister(t *testing.T) {
	t.Skip("Skipping test that requires database connection")

	// Подготавливаем тестовую БД
	database.ConnectDB()

	// Создаем тестовый роутер
	r := SetupTestRouter()
	r.POST("/api/register", controllers.Register)

	// Создаем тестовые данные
	registerInput := models.RegisterInput{
		Username: "testuser",
		Password: "testpassword",
	}
	jsonData, _ := json.Marshal(registerInput)

	// Создаем запрос
	req, _ := http.NewRequest("POST", "/api/register", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	// Выполняем запрос
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Проверяем ответ
	assert.Equal(t, http.StatusCreated, w.Code)

	// Очищаем созданного пользователя
	database.DB.Where("username = ?", "testuser").Delete(&models.User{})
}

// Тест для входа пользователя
func TestLogin(t *testing.T) {
	t.Skip("Skipping test that requires database connection")

	// Подготавливаем тестовую БД
	database.ConnectDB()

	// Создаем тестового пользователя
	r := SetupTestRouter()
	r.POST("/api/register", controllers.Register)

	registerInput := models.RegisterInput{
		Username: "testlogin",
		Password: "testpassword",
	}
	jsonRegisterData, _ := json.Marshal(registerInput)

	req, _ := http.NewRequest("POST", "/api/register", bytes.NewBuffer(jsonRegisterData))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Настраиваем маршрут для входа
	r = SetupTestRouter()
	r.POST("/api/login", controllers.Login)

	// Создаем данные для входа
	loginInput := models.LoginInput{
		Username: "testlogin",
		Password: "testpassword",
	}
	jsonLoginData, _ := json.Marshal(loginInput)

	// Создаем запрос на вход
	req, _ = http.NewRequest("POST", "/api/login", bytes.NewBuffer(jsonLoginData))
	req.Header.Set("Content-Type", "application/json")

	// Выполняем запрос
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Проверяем ответ
	assert.Equal(t, http.StatusOK, w.Code)

	// Проверяем, что в ответе есть токен
	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.NotNil(t, response["token"])

	// Очищаем созданного пользователя
	database.DB.Where("username = ?", "testlogin").Delete(&models.User{})
}

// Тест для попытки входа с неверными учетными данными
func TestLoginWithInvalidCredentials(t *testing.T) {
	t.Skip("Skipping test that requires database connection")

	// Подготавливаем тестовую БД
	database.ConnectDB()

	// Создаем тестовый роутер
	r := SetupTestRouter()
	r.POST("/api/login", controllers.Login)

	// Создаем данные для входа с неверным паролем
	loginInput := models.LoginInput{
		Username: "testlogin",
		Password: "wrongpassword",
	}
	jsonLoginData, _ := json.Marshal(loginInput)

	// Создаем запрос
	req, _ := http.NewRequest("POST", "/api/login", bytes.NewBuffer(jsonLoginData))
	req.Header.Set("Content-Type", "application/json")

	// Выполняем запрос
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Проверяем ответ - должен быть 401 Unauthorized
	assert.Equal(t, http.StatusUnauthorized, w.Code)
}
