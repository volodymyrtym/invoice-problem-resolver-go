package create

import (
	"context"
	"database/sql"
	"errors"
)

type UserCreateRepository struct {
	db  *sql.DB
	ctx context.Context
}

func NewUserCreateRepository(db *sql.DB, ctx context.Context) *UserCreateRepository {
	return &UserCreateRepository{db: db, ctx: ctx}
}

func (r *UserCreateRepository) create(id string, email string, hashedPassword string) error {
	query := "INSERT INTO user_users (id, password, email) VALUES (?, ?, ?)"
	_, err := r.db.ExecContext(r.ctx, query, id, hashedPassword, email)
	if err != nil {
		return err
	}

	return nil
}

func (r *UserCreateRepository) isExists(email string) (bool, error) {
	query := "SELECT 1 FROM user_users WHERE email = ? LIMIT 1"
	var exists int
	err := r.db.QueryRowContext(r.ctx, query, email).Scan(&exists)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) { // Використовуємо errors.Is для перевірки
			return false, nil
		}
		return false, err
	}

	return true, nil
}
