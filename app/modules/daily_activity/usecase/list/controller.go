package list

import (
	"ipr/infra/router/middleware"
	"ipr/infra/template"
	"ipr/modules/daily_activity/authorization"
	"ipr/shared"
	"net/http"
	"strconv"
)

func Controller(handler *Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := authorization.ErrorOnNotAuthorized(w, r, nil)
		if err != nil {
			return
		}

		var query Query

		if r.Method == http.MethodGet {
			if page := r.URL.Query().Get("page"); page != "" {
				if pageInt, err := strconv.Atoi(page); err == nil {
					query.Page = &pageInt
				} else {
					http.Error(w, "Invalid page value", http.StatusBadRequest)
					return
				}
			}

			if rangeStart := r.URL.Query().Get("start_date"); rangeStart != "" {
				query.StartDate = &rangeStart
			}

			if rangeEnd := r.URL.Query().Get("end_date"); rangeEnd != "" {
				query.EndDate = &rangeEnd
			}
		}

		query.UserID = middleware.GetUserIdFromRequest(r)
		data, err := handler.execute(query)
		if err != nil {
			shared.HandleHttpError(w, r, err, nil)
			return
		}

		template.RenderTemplate(w, r, "pages/daily-activities/list.html", data.ToMap())
	}
}
