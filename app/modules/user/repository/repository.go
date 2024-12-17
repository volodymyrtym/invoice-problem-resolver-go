package repository

import (
	"context"
	"database/sql"
	"errors"
)

type UserRepository struct {
	db  *sql.DB
	ctx context.Context
}

func NewUserRepository(db *sql.DB, ctx context.Context) *UserRepository {
	return &UserRepository{db: db, ctx: ctx}
}

func (r *UserRepository) Create(id string, email string, hashedPassword string) error {
	query := "INSERT INTO user_users (id, password, email) VALUES ($1, $2, $3)"
	_, err := r.db.ExecContext(r.ctx, query, id, hashedPassword, email)
	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) isExists(email string) (bool, error) {
	query := "SELECT 1 FROM user_users WHERE email = $1 LIMIT 1"
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
