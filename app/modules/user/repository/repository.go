package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"
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

type User struct {
	Id             string
	Email          string
	HashedPassword string
}

func (r *UserRepository) Find(email string) (*User, error) {
	query := "SELECT id, email, password FROM user_users WHERE email = $1 LIMIT 1"

	var user User
	err := r.db.QueryRowContext(r.ctx, query, email).Scan(&user.Id, &user.Email, &user.HashedPassword)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) IsExists(email string) (bool, error) {
	query := "SELECT 1 FROM user_users WHERE email = $1 LIMIT 1"

	var exists int
	err := r.db.QueryRowContext(r.ctx, query, email).Scan(&exists)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func (r *UserRepository) UpdateLastLogin(userId string) error {
	query := "UPDATE user_users SET last_login_at = $1 WHERE id = $2"

	// Отримання поточного часу
	currentTime := time.Now()

	// Виконання запиту
	_, err := r.db.ExecContext(r.ctx, query, currentTime, userId)
	if err != nil {
		return err
	}

	return nil
}
