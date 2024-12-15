package list

import (
	"fmt"
	"ipr/modules/daily_activity/repository"
	"ipr/shared"
)

type Query struct {
	Page   *int    `json:"page"`
	Range  *string `json:"range"`
	UserID string  `json:"-"`
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
	filter := filterFabric(query)
	activities, _ := handler.repo.GetList(filter)
	views := handler.builder.build(activities.Items)

	return &GetListResult{
		HasNextPage:     activities.HasNextPage,
		HasPreviousPage: activities.HasPreviousPage,
		Items:           views,
	}, nil
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
