package database

import (
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	pgxmigrate "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func RunMigrations() error {
	// Загружаем переменные окружения
	if err := godotenv.Load(); err != nil {
		fmt.Println("Warning: .env file not loaded")
	}

	// Строка подключения
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	if dbHost == "" || dbUser == "" || dbPassword == "" || dbName == "" || dbPort == "" {
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
			dbPort = "5433"
		}
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		dbHost, dbUser, dbPassword, dbName, dbPort,
	)

	// Подключаемся к БД через GORM
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("Failed to connect to database: %v", err)
		return err
	}

	// Получаем соединение SQL
	sqlDB, err := db.DB()
	if err != nil {
		log.Printf("Failed to get database connection: %v", err)
		return err
	}

	// Настраиваем экземпляр migrate
	driver, err := pgxmigrate.WithInstance(sqlDB, &pgxmigrate.Config{})
	if err != nil {
		log.Printf("Failed to create migrate instance: %v", err)
		return err
	}

	// Создаем migrate с источником файлов миграций
	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"postgres", driver)
	if err != nil {
		log.Printf("Failed to create migrate: %v", err)
		return err
	}

	// Запускаем миграции
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Printf("Error running migrations: %v", err)
		// Если ошибка в том, что нет изменений, это не критично
		if err == migrate.ErrNoChange {
			fmt.Println("No migrations to apply")
		} else {
			return err
		}
	} else {
		fmt.Println("Migrations applied successfully")
	}

	// Сохраняем ссылку на DB
	DB = db
	return nil
}
