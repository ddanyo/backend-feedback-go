package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBUser     string
	DBPassword string
	DBHost     string
	DBPort     string
	DBName     string
}

func LoadConfig() (*Config, error) {
	_ = godotenv.Load()

	cfg := &Config{
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBName:     os.Getenv("DB_NAME"),
	}

	if cfg.DBHost == "" {
		cfg.DBHost = "127.0.0.1"
	}
	if cfg.DBPort == "" {
		cfg.DBPort = "5432"
	}

	if err := cfg.Validate(); err != nil {
		return nil, err
	}

	return cfg, nil
}

func (c *Config) Validate() error {
	if c.DBUser == "" {
		return errors.New("Пропущена обязательная переменная окружения: DB_USER")
	}
	if c.DBPassword == "" {
		return errors.New("Пропущена обязательная переменная окружения: DB_PASSWORD")
	}
	if c.DBName == "" {
		return errors.New("Пропущена обязательная переменная окружения: DB_NAME")
	}

	return nil
}
