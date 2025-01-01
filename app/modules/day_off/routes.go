package day_off

import (
	"ipr/infra/router"
	"ipr/modules/day_off/usecase/create_policy"
	"net/http"
)

func RegisterRoutes(deps *Dependencies) {
	router.AddRoute("/api/day-off/policies", http.MethodPost, create_policy.Controller(deps.policyCreateHandler))
}
