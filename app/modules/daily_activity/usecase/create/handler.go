package create

import (
	"errors"
	"fmt"
	"ipr/modules/daily_activity/repository"
	"ipr/shared"
	"time"
)

type Command struct {
	Date         string `json:"date"`
	DateTimeFrom string `json:"dateTimeFrom"`
	DateTimeTo   string `json:"dateTimeTo"`
	Description  string `json:"description"`
	UserID       string `json:"-"`
}

type Handler struct {
	repo *repository.DailyActivityRepository
}

func NewHandler(repo *repository.DailyActivityRepository) *Handler {
	return &Handler{repo: repo}
}

func (handler *Handler) execute(command *Command) (string, error) {
	err := validateRequest(command)
	if err != nil {
		return "", shared.NewInvalidInputError(err.Error())
	}

	date := command.Date
	startAt, err := time.Parse("2006-01-02 15:04", fmt.Sprintf("%s %s", date, command.DateTimeFrom))
	if err != nil {
		return "", shared.NewInvalidInputError("Invalid dateTimeFrom format (expected HH:mm)")
	}
	endAt, err := time.Parse("2006-01-02 15:04", fmt.Sprintf("%s %s", date, command.DateTimeTo))
	if err != nil {
		return "", shared.NewInvalidInputError("Invalid dateTimeTo format (expected HH:mm)")
	}

	id, _ := shared.GenerateGuid()

	activity := repository.CreateDTO{
		ID:          &id,
		UserID:      &command.UserID,
		StartAt:     startAt,
		EndAt:       endAt,
		Description: command.Description,
		CreatedAt:   time.Now(),
		Project:     "Default",
	}

	if err := handler.repo.Save(&activity); err != nil {
		return "", errors.New("failed to save activity")
	}

	return id, nil
}

func validateRequest(c *Command) error {
	if c.Date == "" {
		return fmt.Errorf("date is required")
	}
	if c.DateTimeFrom == "" {
		return fmt.Errorf("dateTimeFrom is required")
	}
	if c.DateTimeTo == "" {
		return fmt.Errorf("dateTimeTo is required")
	}
	if c.Description == "" {
		return fmt.Errorf("description is required")
	}
	if c.UserID == "" {
		return fmt.Errorf("userID is required")
	}

	return nil
}
