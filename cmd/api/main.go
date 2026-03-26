package main

import (
	"log"

	_ "github.com/lib/pq"

	"github.com/ddanyo/backend-feedback-go/internal/config"
	"github.com/ddanyo/backend-feedback-go/internal/db"
	"github.com/ddanyo/backend-feedback-go/internal/feedback"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/ddanyo/backend-feedback-go/docs"
)

// @title           Feedback System API
// @version         1.0
// @description     API сервиса для сбора и отображения отзывов.
// @host            localhost:2510
// @BasePath        /
func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Ошибка инициализации конфига: %v", err)
	}

	db, err := db.Connect(cfg)
	if err != nil {
		log.Fatalf("Ошибка подключения к бд: %v", err)
	}

	repo, err := feedback.NewRepo(db)
	if err != nil {
		log.Fatalf("Ошибка создания репозитория: %v", err)
	}

	handler := feedback.NewHandler(repo)

	router := gin.Default()
	router.GET("/api", handler.GetFeedback)
	router.POST("/api", handler.PostFeedback)

	router.GET("/api/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run(":8080")
}
