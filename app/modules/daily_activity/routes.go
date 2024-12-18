package daily_activity

import (
	"ipr/infra/router"
	"ipr/modules/daily_activity/usecase/create"
	"ipr/modules/daily_activity/usecase/delete"
	"ipr/modules/daily_activity/usecase/list"
	"net/http"
)

func RegisterRoutes(deps *Dependencies) {
	router.AddRoute("/daily-activities/index", http.MethodGet, list.Controller(deps.listHandler))
	router.AddRoute("/daily-activities/index", http.MethodPost, list.Controller(deps.listHandler))
	router.AddRoute("/daily-activities", http.MethodDelete, delete.Controller(deps.repo))
	router.AddRoute("/daily-activities", http.MethodPost, create.Controller(deps.createHandler))
}
