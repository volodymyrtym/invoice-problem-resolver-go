package create_policy

import (
	"errors"
	"fmt"
	"ipr/modules/day_off/repository"
	"ipr/shared"
)

type Command struct {
	Name       string `json:"name"`
	HalfDay    bool   `json:"halfDay"`
	Approvable bool   `json:"approvable"`
	UserID     string `json:"-"`
}

type Handler struct {
	repo *repository.PolicyRepository
}

func NewHandler(repo *repository.PolicyRepository) *Handler {
	return &Handler{repo: repo}
}

func (handler *Handler) execute(command *Command) (string, error) {
	err := validateRequest(command)
	if err != nil {
		return "", shared.NewInvalidInputError(err.Error())
	}

	id, _ := shared.GenerateGuid()
	entity := repository.PolicyCreateDTO{
		Id:         id,
		UserId:     command.UserID,
		Name:       command.Name,
		HalfDay:    command.HalfDay,
		Approvable: command.Approvable,
	}

	if err := handler.repo.SavePolicy(&entity); err != nil {
		return "", errors.New(err.Error())
	}

	return id, nil
}

func validateRequest(c *Command) error {
	if c.Name == "" {
		return fmt.Errorf("name is required")
	}
	if len(c.Name) > 32 {
		return fmt.Errorf("name exceeds maximum length of 32 characters")
	}
	if c.HalfDay != true && c.HalfDay != false {
		return fmt.Errorf("invalid value for HalfDay")
	}
	if c.Approvable != true && c.Approvable != false {
		return fmt.Errorf("invalid value for Approvable")
	}

	return nil
}
