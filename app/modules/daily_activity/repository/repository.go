package repository

import (
	"context"
	"database/sql"
	"errors"
)

type DailyActivityRepository struct {
	db  *sql.DB
	ctx context.Context
}

func NewDailyActivityRepository(db *sql.DB, ctx context.Context) *DailyActivityRepository {
	return &DailyActivityRepository{db: db, ctx: ctx}
}

func (r *DailyActivityRepository) Delete(id string) error {
	query := "DELETE FROM users WHERE id = ? LIMIT 1"
	_, err := r.db.ExecContext(r.ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *DailyActivityRepository) IsOwner(id string, userId string) (bool, error) {
	query := "SELECT 1 FROM user_users WHERE id = ? AND user_id=? LIMIT 1"
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
