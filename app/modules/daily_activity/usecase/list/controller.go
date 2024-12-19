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

		if r.Method == http.MethodPost {
			if err := r.ParseForm(); err != nil {
				http.Error(w, "Invalid form data", http.StatusBadRequest)
				return
			}

			if page := r.FormValue("page"); page != "" {
				if pageInt, err := strconv.Atoi(page); err == nil {
					query.Page = &pageInt
				} else {
					http.Error(w, "Invalid page value", http.StatusBadRequest)
					return
				}
			}

			if rangeVal := r.FormValue("range"); rangeVal != "" {
				query.Range = &rangeVal
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
