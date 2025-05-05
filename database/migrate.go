package database

import (
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	pg "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func RunMigrations() error {
	_ = godotenv.Load()

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	// Подключаемся через GORM, чтобы получить sql.DB
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("❌ не удалось подключиться через GORM: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return fmt.Errorf("❌ ошибка при получении sql.DB: %w", err)
	}

	driver, err := pg.WithInstance(sqlDB, &pg.Config{})
	if err != nil {
		return fmt.Errorf("❌ ошибка инициализации мигратора: %w", err)
	}

	m, err := migrate.NewWithDatabaseInstance("file://migrations", "postgres", driver)
	if err != nil {
		return fmt.Errorf("❌ ошибка создания мигратора: %w", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("❌ ошибка выполнения миграций: %w", err)
	}

	log.Println("✅ SQL миграции успешно применены")
	return nil
}
