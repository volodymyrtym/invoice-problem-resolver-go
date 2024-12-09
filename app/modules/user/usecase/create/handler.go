package create

import (
	"errors"
	"ipr/modules/idgenerator"
	"ipr/modules/user/service/password"
	"net/mail"
)

type UserCreateHandler struct {
	repo              *UserCreateRepository
	passwordValidator *password.Validator
}

func NewUserCreateHandler(repo *UserCreateRepository, passwordValidator *password.Validator) *UserCreateHandler {
	return &UserCreateHandler{repo: repo, passwordValidator: passwordValidator}
}

func (handler *UserCreateHandler) execute(req *createUserRequest) (string, error) {
	if err := handler.validate(req); err != nil {
		return "", err
	}

	hashedPassword, err := password.HashPassword(req.Password)
	if err != nil {
		return "", err
	}

	id, _ := idgenerator.GenerateUserID()
	err = handler.repo.create(id, hashedPassword, req.Email)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (handler *UserCreateHandler) validate(req *createUserRequest) error {
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
