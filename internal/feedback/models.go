package feedback

import "time"

type Feedback struct {
	ID            string    `json:"id" example:"1105"`
	Rating        int       `json:"rating" binding:"required" example:"5"`
	Date_time     time.Time `json:"date_time" example:"2026-03-25T14:30:00Z"`
	Feedback_text string    `json:"feedback_text" example:"Отличный сервис!"`
}

type GetFeedbacksParams struct {
	Skip          int    `form:"skip" example:"0"`
	Take          int    `form:"take" example:"10"`
	Search        string `form:"search" example:"Отличный"`
	SortBy        string `form:"sortBy"`
	CaseSensitive bool   `form:"caseSensitive" example:"false"`
	WholeWord     bool   `form:"wholeWord" example:"false"`
}

type FeedbacksResponse struct {
	Items      []Feedback `json:"items"`
	Total      int        `json:"total" example:"150"`
	TotalPages int        `json:"totalPages" example:"15"`
}

type HTTPError struct {
	Error string `json:"error" example:"Сообщение об ошибке"`
}
