package db

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/ddanyo/backend-feedback-go/internal/config"
	"github.com/joho/godotenv"
)

func Connect(c *config.Config) (*sql.DB, error) {
	_ = godotenv.Load()

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		c.DBHost, c.DBPort, c.DBUser, c.DBPassword, c.DBName)
	if connStr == "" {
		return nil, errors.New("Переменная connStr пустая!")
	}

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("Ошибка создания пула подлючений: %w", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("Ошибка пинга: %w", err)
	}

	return db, nil
}
