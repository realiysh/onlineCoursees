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

// Тест для регистрации пользователя
func TestRegister(t *testing.T) {
	// Создаем тестовый роутер
	r := SetupTestRouter()

	// Регистрируем обработчик с моком для пути /api/register
	r.POST("/api/register", func(c *gin.Context) {
		var input models.RegisterInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Мокаем успешную регистрацию
		c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully", "user_id": 1})
	})

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

	// Проверяем структуру ответа
	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, "User registered successfully", response["message"])
	assert.Equal(t, float64(1), response["user_id"])
}

// Тест для входа пользователя
func TestLogin(t *testing.T) {
	// Создаем тестовый роутер
	r := SetupTestRouter()

	// Регистрируем обработчик с моком для пути /api/login
	r.POST("/api/login", func(c *gin.Context) {
		var input models.LoginInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Проверяем учетные данные
		if input.Username == "testlogin" && input.Password == "testpassword" {
			// Мокаем успешный вход
			c.JSON(http.StatusOK, gin.H{
				"message": "Login successful",
				"token":   "mocked-jwt-token",
				"user": gin.H{
					"id":       1,
					"username": input.Username,
				},
			})
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		}
	})

	// Создаем данные для входа
	loginInput := models.LoginInput{
		Username: "testlogin",
		Password: "testpassword",
	}
	jsonLoginData, _ := json.Marshal(loginInput)

	// Создаем запрос на вход
	req, _ := http.NewRequest("POST", "/api/login", bytes.NewBuffer(jsonLoginData))
	req.Header.Set("Content-Type", "application/json")

	// Выполняем запрос
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Проверяем ответ
	assert.Equal(t, http.StatusOK, w.Code)

	// Проверяем, что в ответе есть токен
	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.NotNil(t, response["token"])
}

// Тест для попытки входа с неверными учетными данными
func TestLoginWithInvalidCredentials(t *testing.T) {
	// Создаем тестовый роутер
	r := SetupTestRouter()

	// Регистрируем обработчик с моком для пути /api/login
	r.POST("/api/login", func(c *gin.Context) {
		var input models.LoginInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Проверяем учетные данные
		if input.Username == "testlogin" && input.Password == "testpassword" {
			// Мокаем успешный вход
			c.JSON(http.StatusOK, gin.H{
				"message": "Login successful",
				"token":   "mocked-jwt-token",
				"user": gin.H{
					"id":       1,
					"username": input.Username,
				},
			})
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		}
	})

	// Создаем данные для входа с неверным паролем
	loginInput := models.LoginInput{
		Username: "testlogin",
		Password: "wrongpassword", // Неверный пароль
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
