package database

import (
	"fmt"
	"log"
	"os"
	"user-service/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() error {
	// Загружаем переменные окружения
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️  .env файл не найден, продолжаем с переменными окружения")
	}

	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		dbHost, dbUser, dbPassword, dbName, dbPort,
	)

	// Подключаемся к БД через GORM
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("❌ Ошибка подключения к базе данных: %v", err)
		return err
	}

	// Принудительно создаем таблицы
	if err := db.Migrator().DropTable(&models.User{}, &models.Author{}); err != nil {
		log.Printf("❌ Ошибка удаления таблиц: %v", err)
	}

	// Выполняем автоматическую миграцию
	if err := db.AutoMigrate(&models.User{}, &models.Author{}); err != nil {
		log.Printf("❌ AutoMigrate ошибка: %v", err)
		return err
	}

	// Сохраняем подключение
	DB = db
	log.Println("✅ Подключение к базе установлено и миграции выполнены")
	return nil
}
