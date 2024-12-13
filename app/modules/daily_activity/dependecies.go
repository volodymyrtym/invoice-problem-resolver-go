package daily_activity

import (
	"context"
	"database/sql"
	"ipr/modules/daily_activity/authorization"
	"ipr/modules/daily_activity/repository"
)

type Dependencies struct {
	repo *repository.DailyActivityRepository
	auth *authorization.Auth
}

func NewDependencies(db *sql.DB, ctx context.Context) *Dependencies {
	repo := repository.NewDailyActivityRepository(db, ctx)

	return &Dependencies{
		repo: repo,
		auth: authorization.NewAuth(repo),
	}
}
