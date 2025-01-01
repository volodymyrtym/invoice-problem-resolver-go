package list

import (
	"errors"
	"fmt"
	"ipr/modules/daily_activity/repository"
	"ipr/shared"
	"time"
)

var limitDays = 31

type Query struct {
	Page      *int    `json:"page"`
	StartDate *string `json:"startDate"`
	EndDate   *string `json:"endDate"`
	UserID    string  `json:"-"`
}

type GetListResult struct {
	HasPreviousPage bool                        `json:"hasPreviousPage"`
	HasNextPage     bool                        `json:"hasNextPage"`
	Items           []DailyActivitiesCollection `json:"items"`
}

type Handler struct {
	repo    *repository.DailyActivityRepository
	builder *ResultItemsBuilder
}

func NewHandler(repo *repository.DailyActivityRepository, b *ResultItemsBuilder) *Handler {
	return &Handler{repo: repo, builder: b}
}

func (handler *Handler) execute(query Query) (*GetListResult, error) {
	err := validateRequest(query)
	if err != nil {
		return nil, shared.NewInvalidInputError(err.Error())
	}

	filter := repository.GetListFilter{
		UserID: query.UserID,
	}

	if query.StartDate != nil {
		startAt, err := parseDate(query.StartDate)
		if err != nil {
			return nil, err
		}
		filter.StartDate = startAt
	}
	if query.EndDate != nil {
		endAt, err := parseDate(query.EndDate)
		if err != nil {
			return nil, err
		}
		filter.EndDate = endAt
	}

	if query.StartDate == nil && query.EndDate == nil {
		filter.Page = query.Page
		filter.Limit = &limitDays
	}

	activities, err := handler.repo.GetList(&filter)
	if err != nil {
		return nil, err
	}

	views := handler.builder.build(activities.Items)

	return &GetListResult{
		HasNextPage:     activities.HasNextPage,
		HasPreviousPage: activities.HasPreviousPage,
		Items:           views,
	}, nil
}
func parseDate(dateStr *string) (*time.Time, error) {
	if dateStr == nil || *dateStr == "" {
		return nil, errors.New("date is required")
	}

	parsedDate, err := time.Parse("2006-01-02", *dateStr)
	if err != nil {
		return nil, fmt.Errorf("cannot parse date %s: %w", *dateStr, err)
	}
	return &parsedDate, nil
}

func validateRequest(c Query) error {
	if c.UserID == "" {
		return fmt.Errorf("userID is required")
	}

	return nil
}

func (r GetListResult) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"HasNextPage":     r.HasNextPage,
		"HasPreviousPage": r.HasPreviousPage,
		"Items":           r.Items,
	}
}
