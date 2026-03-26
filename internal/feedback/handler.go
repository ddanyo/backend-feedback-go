package feedback

import (
	"fmt"
	"math"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type FeedbackService interface {
	SelectFeedbacks(par GetFeedbacksParams) ([]Feedback, int, error)
	AddFeedback(fd Feedback) (string, error)
}

type Handler struct {
	fservice FeedbackService
}

func NewHandler(f FeedbackService) *Handler {
	return &Handler{
		fservice: f,
	}
}

// GetFeedback godoc
// @Summary      Получить список отзывов
// @Description  Возвращает список отзывов с поддержкой пагинации, поиска и сортировки
// @Tags         Feedbacks
// @Produce      json
// @Param        skip          query int    false "Сколько записей пропустить (offset)" default(0)
// @Param        take          query int    false "Сколько записей вернуть (limit)" default(10)
// @Param        search        query string false "Поисковый запрос"
// @Param        caseSensitive query bool   false "Учитывать регистр при поиске" default(false)
// @Param        wholeWord     query bool   false "Искать только целые слова" default(false)
// @Success      200 {object}  FeedbacksResponse
// @Failure      400 {object} HTTPError "Неверные параметры запроса"
// @Failure      500 {object} HTTPError "Внутренняя ошибка сервера"
// @Router       /api [get]
func (h *Handler) GetFeedback(c *gin.Context) {

	var params GetFeedbacksParams
	if err := c.ShouldBindQuery(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверные параметры запроса"})
		return
	}

	if params.Skip < 0 {
		params.Skip = 0
	}
	if params.Take < 5 {
		params.Take = 5
	}
	if params.Take > 100 {
		params.Take = 100
	}
	if params.SortBy == "" {
		params.SortBy = "newest"
	}

	items, total, err := h.fservice.SelectFeedbacks(params)
	if err != nil {
		fmt.Printf("[ERROR] Ошибка получения отзывов: %v\n", err)
		// gin.H = map[string]any
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"error": "Не удалось загрузить отзывы. Попробуйте позже.",
		})
		return
	}

	totalPages := int(math.Ceil(float64(total) / float64(params.Take)))
	if totalPages == 0 {
		totalPages = 1
	}

	response := FeedbacksResponse{
		Items:      items,
		Total:      total,
		TotalPages: totalPages,
	}

	c.IndentedJSON(http.StatusOK, response)
}

// PostFeedback godoc
// @Summary      Добавить новый отзыв
// @Description  Создает новый отзыв в базе данных
// @Tags         Feedbacks
// @Accept       json
// @Produce      json
// @Param        data body     Feedback true "Данные отзыва"
// @Success      201      {object} Feedback
// @Failure      400 {object} HTTPError "Неверные параметры запроса"
// @Failure      500 {object} HTTPError "Внутренняя ошибка сервера"
// @Router       /api [post]
func (h *Handler) PostFeedback(c *gin.Context) {
	var newFeed Feedback

	if err := c.ShouldBindJSON(&newFeed); err != nil {
		fmt.Printf("[WARN] Некорректный JSON: %v\n", err)

		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": "Неверный формат данных",
		})
		return
	}

	newFeed.Date_time = time.Now()

	id, err := h.fservice.AddFeedback(newFeed)
	if err != nil {
		fmt.Printf("[ERROR] Ошибка сохранения отзыва в БД: %v\n", err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"error": "Не удалось сохранить отзыв",
		})
		return
	}

	newFeed.ID = id
	c.IndentedJSON(http.StatusCreated, newFeed)
}
