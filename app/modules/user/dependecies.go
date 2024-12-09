package user

import (
	"context"
	"database/sql"
	"ipr/modules/user/service/password"
	"ipr/modules/user/usecase/create"
)

type Dependencies struct {
	Repo              *create.UserCreateRepository
	PasswordValidator *password.Validator
	Handler           *create.UserCreateHandler
}

func NewDependencies(db *sql.DB, ctx context.Context) *Dependencies {
	passwordValidator := password.NewPasswordValidator()
	repo := create.NewUserCreateRepository(db, ctx)
	handler := create.NewUserCreateHandler(repo, passwordValidator)

	return &Dependencies{
		Repo:              repo,
		PasswordValidator: passwordValidator,
		Handler:           handler,
	}
}
