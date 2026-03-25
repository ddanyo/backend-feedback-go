package feedback

import "time"

type Feedback struct {
	ID            string    `json:"id"`
	Rating        int       `json:"rating"`
	Date_time     time.Time `json:"date_time"`
	Feedback_text string    `json:"feedback_text"`
}

type GetFeedbacksParams struct {
	Skip          int    `form:"skip"`
	Take          int    `form:"take"`
	Search        string `form:"search"`
	SortBy        string `form:"sortBy"`
	CaseSensitive bool   `form:"caseSensitive"`
	WholeWord     bool   `form:"wholeWord"`
}

type FeedbacksResponse struct {
	Items      []Feedback `json:"items"`
	Total      int        `json:"total"`
	TotalPages int        `json:"totalPages"`
}
