package database

import (
	"fmt"
	"os"
	"project1/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {

	if err := godotenv.Load(); err != nil {
		fmt.Println("Warning: .env file not loaded")
	}

	// Строка подключения
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	// Открытие соединения с базой данных
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	// Автоматическая миграция
	if err := db.AutoMigrate(&models.User{}, &models.Course{}, &models.Author{}); err != nil {
		panic("AutoMigrate failed")
	}

	// Сохраняем ссылку на DB
	DB = db
}
