package login

import (
	"context"
	"database/sql"
	"errors"
)

type UserLoginRepository struct {
	db  *sql.DB
	ctx context.Context
}

func NewUserLoginRepository(db *sql.DB, ctx context.Context) *UserLoginRepository {
	return &UserLoginRepository{db: db, ctx: ctx}
}

type User struct {
	Email          string
	HashedPassword string
	Id             string
}

func (r *UserLoginRepository) find(email string) (*User, error) {
	query := "SELECT id, email, password FROM user_users WHERE email = ? LIMIT 1"

	var user User
	err := r.db.QueryRowContext(r.ctx, query, email).Scan(&user.Email, &user.HashedPassword)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}
