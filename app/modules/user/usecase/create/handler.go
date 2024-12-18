package create

import (
	"errors"
	"ipr/modules/user/repository"
	"ipr/modules/user/service/password"
	"ipr/shared"
	"net/mail"
)

type UserCreateHandler struct {
	repo              *repository.UserRepository
	passwordValidator *password.Validator
}

type command struct {
	Email    string `json:"Email"`
	Password string `json:"Password"`
}

func NewUserCreateHandler(repo *repository.UserRepository, passwordValidator *password.Validator) *UserCreateHandler {
	return &UserCreateHandler{repo: repo, passwordValidator: passwordValidator}
}

func (handler *UserCreateHandler) execute(cmd *command) (string, error) {
	if err := handler.validate(cmd); err != nil {
		return "", shared.NewInvalidInputError(err.Error())
	}

	isExists, _ := handler.repo.IsExists(cmd.Email)
	if isExists {
		return "", shared.NewInvalidInputError("user already exists")
	}

	hashedPassword, err := password.HashPassword(cmd.Password)
	if err != nil {
		return "", err
	}

	id, _ := shared.GenerateGuid()
	err = handler.repo.Create(id, cmd.Email, hashedPassword)
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
