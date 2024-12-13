package daily_activity

import (
	"ipr/infra/router"
	"ipr/infra/session"
	"ipr/modules/daily-activity/usecase/list"
	"net/http"
)

func RegisterRoutes(deps *Dependencies, sm *session.SessionManager) {
	router.AddRoute("/daily-activities", http.MethodGet, list.Controller())
	router.AddRoute("/daily-activities", http.MethodDelete, func(writer http.ResponseWriter, request *http.Request) {

	})
	router.AddRoute("/daily-activities", http.MethodPost, func(writer http.ResponseWriter, request *http.Request) {

	})
}
