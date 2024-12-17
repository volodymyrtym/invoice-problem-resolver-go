package list

import (
	"encoding/json"
	"ipr/infra/router/middleware"
	"ipr/infra/template"
	"ipr/modules/daily_activity/authorization"
	"ipr/shared"
	"net/http"
)

func Controller(handler *Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := authorization.ErrorOnNotAuthorized(w, r, nil)
		if err != nil {
			return
		}

		var query Query
		if err := json.NewDecoder(r.Body).Decode(&query); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		query.UserID = middleware.GetUserIdFromRequest(r)
		data, err := handler.execute(query)
		if err != nil {
			shared.HandleHttpError(w, r, err, nil)
			return
		}

		template.RenderTemplate(w, "pages/daily-activities/daily-activities-list.html", data.ToMap())
	}
}
