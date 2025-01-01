package day_off

import (
	"context"
	"database/sql"
	auth "ipr/modules/day_off/authorization"
	repositories "ipr/modules/day_off/repository"
	"ipr/modules/day_off/usecase/create_policy"
)

type Dependencies struct {
	policyRepo          *repositories.PolicyRepository
	policyCreateHandler *create_policy.Handler
}

func NewDependencies(db *sql.DB, ctx context.Context) *Dependencies {
	policyRepository := repositories.NewPolicyRepository(db, ctx)
	auth.Initialize(policyRepository)

	return &Dependencies{
		policyRepo:          policyRepository,
		policyCreateHandler: create_policy.NewHandler(policyRepository),
	}
}
