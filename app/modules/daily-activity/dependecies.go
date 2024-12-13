package daily_activity

import (
	"context"
	"database/sql"
	"ipr/modules/user/service/password"
	"ipr/modules/user/usecase/create"
	"ipr/modules/user/usecase/login"
)

type Dependencies struct {
	CreateHandler *create.UserCreateHandler
	LoginHandler  *login.UserLoginHandler
}

func NewDependencies(db *sql.DB, ctx context.Context) *Dependencies {
	createPasswordValidator := password.NewPasswordValidator()

	createRepo := create.NewUserCreateRepository(db, ctx)
	createHandler := create.NewUserCreateHandler(createRepo, createPasswordValidator)

	loginRepo := login.NewUserLoginRepository(db, ctx)
	loginHandler := login.NewUserLoginHandler(loginRepo, createPasswordValidator)

	return &Dependencies{
		CreateHandler: createHandler,
		LoginHandler:  loginHandler,
	}
}
