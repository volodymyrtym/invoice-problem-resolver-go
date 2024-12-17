package login

import (
	"ipr/modules/user/repository"
	"ipr/modules/user/service/password"
	"ipr/shared"
)

type UserLoginHandler struct {
	repo              *repository.UserRepository
	passwordValidator *password.Validator
}

type command struct {
	Email    string `json:"Email"`
	Password string `json:"Password"`
}

func NewUserLoginHandler(repo *repository.UserRepository, passwordValidator *password.Validator) *UserLoginHandler {
	return &UserLoginHandler{repo: repo, passwordValidator: passwordValidator}
}

func (handler *UserLoginHandler) execute(command *command) (string, error) {
	user, _ := handler.repo.Find(command.Email)
	if user == nil {
		return "", shared.NewInvalidInputError("unknown email")
	}

	if password.VerifyPassword(command.Password, user.HashedPassword) == false {
		return "", shared.NewInvalidInputError("wrong password")
	}

	handler.repo.UpdateLastLogin(user.Id)

	return user.Id, nil
}
