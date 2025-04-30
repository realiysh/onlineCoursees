package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() error {
	// Загружаем переменные окружения
	if err := godotenv.Load(); err != nil {
		fmt.Println("Warning: .env file not loaded")
	}

	// Проверяем наличие переменных окружения
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	if dbHost == "" || dbUser == "" || dbPassword == "" || dbName == "" || dbPort == "" {
		log.Println("Database connection details are missing in the .env file")
		log.Printf("Host: %s, User: %s, DB: %s, Port: %s\n", dbHost, dbUser, dbName, dbPort)
		// Устанавливаем значения по умолчанию, если не найдены в .env
		if dbHost == "" {
			dbHost = "localhost"
		}
		if dbUser == "" {
			dbUser = "postgres"
		}
		if dbPassword == "" {
			dbPassword = "postgres"
		}
		if dbName == "" {
			dbName = "postgres"
		}
		if dbPort == "" {
			dbPort = "5432"
		}
	}

	// Строка подключения
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		dbHost, dbUser, dbPassword, dbName, dbPort,
	)

	log.Printf("Connecting to database at %s:%s with user %s", dbHost, dbPort, dbUser)

	// Открытие соединения с базой данных
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("failed to initialize database, got error %v", err)
		return err
	}

	log.Println("Database connection established successfully")

	// Запускаем миграции
	return RunMigrations()
}
