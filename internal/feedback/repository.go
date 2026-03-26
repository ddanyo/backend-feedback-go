package feedback

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"
)

type postgresRepo struct {
	db *sql.DB
}

func NewRepo(database *sql.DB) (*postgresRepo, error) {
	if database == nil {
		return nil, errors.New("Пустой дескриптор базы данных")
	}

	return &postgresRepo{
		db: database,
	}, nil
}

func (p *postgresRepo) SelectFeedbacks(par GetFeedbacksParams) ([]Feedback, int, error) {

	var conditions []string
	var args []any
	argID := 1

	conditions = append(conditions, "1=1")

	if par.Search != "" {
		if par.WholeWord {
			operator := "~"
			if !par.CaseSensitive {
				operator = "~*"
			}

			conditions = append(conditions, fmt.Sprintf("feedback_text %s $%d", operator, argID))
			args = append(args, `\y`+par.Search+`\y`)
			argID++
		} else {
			if par.CaseSensitive {
				conditions = append(conditions, fmt.Sprintf("feedback_text LIKE $%d", argID))
			} else {
				conditions = append(conditions, fmt.Sprintf("feedback_text ILIKE $%d", argID))
			}
			args = append(args, "%"+par.Search+"%")
			argID++
		}
	}

	whereClause := strings.Join(conditions, " AND ")

	var total int
	countQuery := fmt.Sprintf("SELECT COUNT(id) FROM feedbacks_table WHERE %s", whereClause)

	err := p.db.QueryRow(countQuery, args...).Scan(&total)
	if err != nil {
		return nil, 0, fmt.Errorf("ошибка подсчета: %v", err)
	}

	if total == 0 {
		return []Feedback{}, 0, nil
	}

	orderBy := "ORDER BY id DESC, date_time DESC"
	switch par.SortBy {
	case "oldest":
		orderBy = "ORDER BY id ASC, date_time ASC"
	case "rating_high":
		orderBy = "ORDER BY rating DESC, id DESC"
	case "rating_low":
		orderBy = "ORDER BY rating ASC, id DESC"
	}

	limitOffset := fmt.Sprintf("LIMIT $%d OFFSET $%d", argID, argID+1)
	args = append(args, par.Take, par.Skip)

	finalQuery := fmt.Sprintf("SELECT id, rating, date_time, feedback_text FROM feedbacks_table WHERE %s %s %s", whereClause, orderBy, limitOffset)

	rows, err := p.db.Query(finalQuery, args...)
	if err != nil {
		return nil, 0, fmt.Errorf("Ошибка получения данных из бд: %w", err)
	}
	defer rows.Close()

	var feedbacks []Feedback

	for rows.Next() {
		var fd Feedback
		if err := rows.Scan(&fd.ID, &fd.Rating, &fd.Date_time, &fd.Feedback_text); err != nil {
			return nil, 0, fmt.Errorf("Ошибка при сканировании строки отзыва: %w", err)
		}
		feedbacks = append(feedbacks, fd)
	}
	if err := rows.Err(); err != nil {
		return nil, 0, fmt.Errorf("Ошибка при переборе строк данных: %w", err)
	}
	return feedbacks, total, nil
}

func (p *postgresRepo) AddFeedback(fd Feedback) (string, error) {
	var id string

	query := "INSERT INTO feedbacks_table (rating, date_time, feedback_text) VALUES ($1, $2, $3) RETURNING id"

	err := p.db.QueryRow(query, fd.Rating, fd.Date_time, fd.Feedback_text).Scan(&id)
	if err != nil {
		return "", fmt.Errorf("Ошибка добавления данных в бд: %w", err)
	}

	return id, nil
}
