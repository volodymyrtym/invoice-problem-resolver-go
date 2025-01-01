package user

import (
	"context"
	"database/sql"
	"ipr/modules/user/repository"
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

	userRepository := repository.NewUserRepository(db, ctx)
	createHandler := create.NewUserCreateHandler(userRepository, createPasswordValidator)
	loginHandler := login.NewUserLoginHandler(userRepository, createPasswordValidator)

	return &Dependencies{
		CreateHandler: createHandler,
		LoginHandler:  loginHandler,
	}
}
