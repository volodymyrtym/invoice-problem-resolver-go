package daily_activity

import (
	"context"
	"database/sql"
	"ipr/modules/daily_activity/authorization"
	"ipr/modules/daily_activity/repository"
	"ipr/modules/daily_activity/usecase/create"
	"ipr/modules/daily_activity/usecase/list"
)

type Dependencies struct {
	repo          *repository.DailyActivityRepository
	createHandler *create.Handler
	listHandler   *list.Handler
}

func NewDependencies(db *sql.DB, ctx context.Context) *Dependencies {
	repo := repository.NewDailyActivityRepository(db, ctx)
	authorization.NewAuth(repo)
	createHandler := create.NewHandler(repo)
	listHandler := list.NewHandler(repo, list.NewResultItemsBuilder())

	return &Dependencies{
		repo:          repo,
		createHandler: createHandler,
		listHandler:   listHandler,
	}
}
