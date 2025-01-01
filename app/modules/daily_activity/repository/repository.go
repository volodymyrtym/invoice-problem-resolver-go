package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"
)

type DailyActivityRepository struct {
	db  *sql.DB
	ctx context.Context
}

func NewDailyActivityRepository(db *sql.DB, ctx context.Context) *DailyActivityRepository {
	return &DailyActivityRepository{db: db, ctx: ctx}
}

func (r *DailyActivityRepository) Delete(id string) error {
	query := "DELETE FROM daily_activity_daily_activities WHERE id = $1"
	_, err := r.db.ExecContext(r.ctx, query, id)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

type CreateDTO struct {
	ID          *string
	UserID      *string
	StartAt     time.Time
	EndAt       time.Time
	Description string
	CreatedAt   time.Time
	Project     string
}

func (r *DailyActivityRepository) Save(dto *CreateDTO) error {
	query := `
		INSERT INTO daily_activity_daily_activities 
		(id, user_id, start_at, end_at, description, created_at, project) 
		VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err := r.db.Exec(query, dto.ID, dto.UserID, dto.StartAt, dto.EndAt, dto.Description, dto.CreatedAt, dto.Project)
	return err
}

type GetListFilter struct {
	Page      *int
	UserID    string
	Limit     *int
	StartDate *time.Time
	EndDate   *time.Time
}

type ListQueryResult struct {
	HasPreviousPage bool
	HasNextPage     bool
	Items           []QueryItem
}

type QueryItem struct {
	ID          string
	ProjectName string
	StartAt     time.Time
	EndAt       time.Time
	Description string
}

func (r *DailyActivityRepository) GetList(filter *GetListFilter) (*ListQueryResult, error) {
	var whereClauses []string
	var args []interface{}
	args = append(args, filter.UserID)
	whereClauses = append(whereClauses, "user_id = $1")

	if filter.StartDate != nil && filter.EndDate != nil {
		args = append(args, filter.StartDate, filter.EndDate)
		whereClauses = append(whereClauses, "DATE(start_at) BETWEEN $2 AND $3")
	} else if filter.StartDate != nil {
		args = append(args, filter.StartDate)
		whereClauses = append(whereClauses, "DATE(start_at) >= $2")
	} else if filter.EndDate != nil {
		args = append(args, filter.EndDate)
		whereClauses = append(whereClauses, "DATE(start_at) <= $2")
	}

	sqlQuery := fmt.Sprintf(`
		SELECT id, project AS project_name, start_at, end_at, description
		FROM daily_activity_daily_activities
		WHERE %s
		ORDER BY start_at DESC
	`, joinWhereClauses(whereClauses))

	fmt.Println("sqlQuery: ", sqlQuery)
	fmt.Println("args: ", args)
	if filter.Page != nil && filter.Limit != nil {
		offset := (*filter.Page - 1) * *filter.Limit
		sqlQuery += fmt.Sprintf(" LIMIT %d OFFSET %d", filter.Limit, offset)
	}

	rows, err := r.db.QueryContext(r.ctx, sqlQuery, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []QueryItem
	for rows.Next() {
		var item QueryItem
		var startAt, endAt time.Time
		var projectName sql.NullString
		if err := rows.Scan(&item.ID, &projectName, &startAt, &endAt, &item.Description); err != nil {
			log.Println("error", err)
			return nil, err
		}
		item.ProjectName = projectName.String
		item.StartAt = startAt
		item.EndAt = endAt
		items = append(items, item)
	}

	hasNextPage := false
	if filter.Page != nil && len(items) > 0 {
		oldestDate := items[len(items)-1].StartAt
		hasNextPage, err = r.hasNextDate(filter.UserID, oldestDate)
		if err != nil {
			return nil, err
		}
	}

	return &ListQueryResult{
		HasPreviousPage: filter.Page != nil && *filter.Page > 1,
		HasNextPage:     hasNextPage,
		Items:           items,
	}, nil
}

func (r *DailyActivityRepository) hasNextDate(userID string, oldestDate time.Time) (bool, error) {
	var args []interface{}
	args = append(args, userID, oldestDate.Format("2006-01-02"))

	whereClauses := []string{
		"user_id = $1",
		"DATE(start_at) > $2",
	}

	sqlQuery := fmt.Sprintf(`
		SELECT EXISTS (
			SELECT 1
			FROM daily_activity_daily_activities
			WHERE %s
		) AS next_exists
	`, joinWhereClauses(whereClauses))

	var exists bool
	err := r.db.QueryRowContext(r.ctx, sqlQuery, args...).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}

func joinWhereClauses(clauses []string) string {
	return strings.Join(clauses, " AND ")
}

func (r *DailyActivityRepository) IsOwner(id string, userId string) (bool, error) {
	query := "SELECT 1 FROM daily_activity_daily_activities WHERE id = $1 AND user_id = $2 LIMIT 1"
	var exists int
	err := r.db.QueryRowContext(r.ctx, query, id, userId).Scan(&exists)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) { // Використовуємо errors.Is для перевірки
			return false, nil
		}
		return false, err
	}

	return true, nil
}
