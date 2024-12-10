package login

import (
	"ipr/common"
	"ipr/modules/user/service/password"
)

type UserLoginHandler struct {
	repo              *UserLoginRepository
	passwordValidator *password.Validator
}

func NewUserLoginHandler(repo *UserLoginRepository, passwordValidator *password.Validator) *UserLoginHandler {
	return &UserLoginHandler{repo: repo, passwordValidator: passwordValidator}
}

func (handler *UserLoginHandler) execute(command *createCommand) (string, error) {
	user, _ := handler.repo.find(command.Email)
	if user == nil {
		return "", common.NewInvalidInputError("unknown email")
	}

	if password.VerifyPassword(command.Password, user.HashedPassword) == false {
		return "", common.NewInvalidInputError("wrong password")
	}

	return user.Id, nil
}
