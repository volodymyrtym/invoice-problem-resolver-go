package create

import (
	"errors"
	"ipr/modules/user/service/password"
	"ipr/shared"
	"net/mail"
)

type UserCreateHandler struct {
	repo              *UserCreateRepository
	passwordValidator *password.Validator
}

type command struct {
	Email    string
	Password string
}

func NewUserCreateHandler(repo *UserCreateRepository, passwordValidator *password.Validator) *UserCreateHandler {
	return &UserCreateHandler{repo: repo, passwordValidator: passwordValidator}
}

func (handler *UserCreateHandler) execute(req *command) (string, error) {
	if err := handler.validate(req); err != nil {
		return "", shared.NewInvalidInputError(err.Error())
	}

	hashedPassword, err := password.HashPassword(req.Password)
	if err != nil {
		return "", err
	}

	id, _ := shared.GenerateGuid()
	err = handler.repo.create(id, hashedPassword, req.Email)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (handler *UserCreateHandler) validate(req *command) error {
	err := handler.passwordValidator.Validate(req.Password)
	if err != nil {
		return err
	}

	if isValidEmail(req.Email) == false {
		return errors.New("invalid Email")
	}

	return nil
}

func isValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)

	return err == nil
}
