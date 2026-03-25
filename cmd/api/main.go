package main

import (
	"log"

	_ "github.com/lib/pq"

	"github.com/ddanyo/backend-feedback-go/internal/config"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Ошибка инициализации конфига: %v", err)
	}

}
