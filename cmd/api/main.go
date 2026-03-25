package main

import (
	"log"

	_ "github.com/lib/pq"

	"github.com/ddanyo/backend-feedback-go/internal/config"
	"github.com/ddanyo/backend-feedback-go/internal/db"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Ошибка инициализации конфига: %v", err)
	}

	db, err := db.Connect(cfg)
	if err != nil {
		log.Fatalf("Ошибка подключения к бд: %v", err)
	}
}
