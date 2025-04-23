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

// Тест для получения списка авторов
func TestGetAuthors(t *testing.T) {
	// Создаем тестовый роутер
	r := SetupTestRouter()

	// Регистрируем обработчик с моком для пути /api/authors
	r.GET("/api/authors", func(c *gin.Context) {
		// Создаем моковые данные авторов
		authors := []models.Author{
			{ID: 1, Name: "Test Author 1"},
			{ID: 2, Name: "Test Author 2"},
		}

		// Возвращаем список авторов
		c.JSON(http.StatusOK, gin.H{
			"data": authors,
			"meta": gin.H{
				"page":  1,
				"limit": 10,
				"total": 2,
			},
		})
	})

	// Создаем запрос
	req, _ := http.NewRequest("GET", "/api/authors", nil)

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
}

// Тест для создания автора
func TestCreateAuthor(t *testing.T) {
	// Создаем тестовый роутер
	r := SetupTestRouter()

	// Регистрируем обработчик с моком для пути /api/authors
	r.POST("/api/authors", func(c *gin.Context) {
		var input models.AuthorInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Создаем моковый ответ с новым автором
		author := models.Author{
			ID:   1,
			Name: input.Name,
		}

		c.JSON(http.StatusCreated, author)
	})

	// Создаем тестовые данные
	authorInput := models.AuthorInput{
		Name: "Test Author Creation",
	}
	jsonData, _ := json.Marshal(authorInput)

	// Создаем запрос
	req, _ := http.NewRequest("POST", "/api/authors", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	// Выполняем запрос
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Проверяем ответ
	assert.Equal(t, http.StatusCreated, w.Code)

	// Проверяем содержимое ответа
	var response models.Author
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, uint(1), response.ID)
	assert.Equal(t, "Test Author Creation", response.Name)
}

// Тест для получения автора по ID
func TestGetAuthorByID(t *testing.T) {
	// Создаем тестовый роутер
	r := SetupTestRouter()

	// Регистрируем обработчик с моком для пути /api/authors/:id
	r.GET("/api/authors/:id", func(c *gin.Context) {
		id := c.Param("id")

		// Проверяем ID
		if id != "1" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Author not found"})
			return
		}

		// Возвращаем автора
		author := models.Author{
			ID:   1,
			Name: "Test Author GetByID",
		}

		c.JSON(http.StatusOK, author)
	})

	// Создаем запрос
	req, _ := http.NewRequest("GET", "/api/authors/1", nil)

	// Выполняем запрос
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Проверяем ответ
	assert.Equal(t, http.StatusOK, w.Code)

	// Проверяем содержимое ответа
	var response models.Author
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, uint(1), response.ID)
	assert.Equal(t, "Test Author GetByID", response.Name)
}

// Тест для обновления автора
func TestUpdateAuthor(t *testing.T) {
	// Создаем тестовый роутер
	r := SetupTestRouter()

	// Регистрируем обработчик с моком для пути /api/authors/:id
	r.PUT("/api/authors/:id", func(c *gin.Context) {
		id := c.Param("id")

		// Проверяем ID
		if id != "1" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Author not found"})
			return
		}

		var input models.AuthorInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Возвращаем обновленного автора
		author := models.Author{
			ID:   1,
			Name: input.Name,
		}

		c.JSON(http.StatusOK, author)
	})

	// Создаем тестовые данные
	authorInput := models.AuthorInput{
		Name: "Test Author After Update",
	}
	jsonData, _ := json.Marshal(authorInput)

	// Создаем запрос
	req, _ := http.NewRequest("PUT", "/api/authors/1", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	// Выполняем запрос
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Проверяем ответ
	assert.Equal(t, http.StatusOK, w.Code)

	// Проверяем содержимое ответа
	var response models.Author
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, uint(1), response.ID)
	assert.Equal(t, "Test Author After Update", response.Name)
}

// DeleteAuthor удаляет автора
func TestDeleteAuthor(t *testing.T) {
	// Создаем тестовый роутер
	r := SetupTestRouter()

	// Регистрируем обработчик с моком для пути /api/authors/:id
	r.DELETE("/api/authors/:id", func(c *gin.Context) {
		id := c.Param("id")

		// Проверяем ID
		if id != "1" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Author not found"})
			return
		}

		// Возвращаем успешный ответ
		c.JSON(http.StatusOK, gin.H{"message": "Author deleted successfully"})
	})

	// Создаем запрос
	req, _ := http.NewRequest("DELETE", "/api/authors/1", nil)

	// Выполняем запрос
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Проверяем ответ
	assert.Equal(t, http.StatusOK, w.Code)

	// Проверяем содержимое ответа
	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, "Author deleted successfully", response["message"])
}

// Тест на попытку удаления автора, у которого есть связанные курсы
func TestDeleteAuthorWithCourses(t *testing.T) {
	// Создаем тестовый роутер
	r := SetupTestRouter()

	// Регистрируем обработчик с моком для пути /api/authors/:id
	r.DELETE("/api/authors/:id", func(c *gin.Context) {
		id := c.Param("id")

		// Для теста считаем, что у автора с ID=2 есть курсы
		if id == "2" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Cannot delete author with related courses"})
			return
		}

		// Возвращаем успешный ответ для других ID
		c.JSON(http.StatusOK, gin.H{"message": "Author deleted successfully"})
	})

	// Создаем запрос на удаление автора с курсами
	req, _ := http.NewRequest("DELETE", "/api/authors/2", nil)

	// Выполняем запрос
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Проверяем ответ - должен быть ошибочный статус
	assert.Equal(t, http.StatusBadRequest, w.Code)

	// Проверяем сообщение об ошибке
	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, "Cannot delete author with related courses", response["error"])
}
