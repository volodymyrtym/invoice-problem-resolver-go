package login

import (
	"ipr/modules/user/service/password"
	"ipr/shared"
)

type UserLoginHandler struct {
	repo              *UserLoginRepository
	passwordValidator *password.Validator
}

type command struct {
	Email    string `json:"Email"`
	Password string `json:"HashedPassword"`
}

func NewUserLoginHandler(repo *UserLoginRepository, passwordValidator *password.Validator) *UserLoginHandler {
	return &UserLoginHandler{repo: repo, passwordValidator: passwordValidator}
}

func (handler *UserLoginHandler) execute(command *command) (string, error) {
	user, _ := handler.repo.find(command.Email)
	if user == nil {
		return "", shared.NewInvalidInputError("unknown email")
	}

	if password.VerifyPassword(command.Password, user.HashedPassword) == false {
		return "", shared.NewInvalidInputError("wrong password")
	}

	return user.Id, nil
}
