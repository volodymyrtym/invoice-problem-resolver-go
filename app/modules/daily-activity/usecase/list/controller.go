package list

import (
	"ipr/infra/template"
	"net/http"
)

func Controller() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := map[string]interface{}{
			"message": "Welcome to Invoice Problem Resolver",
		}
		template.RenderTemplate(w, "pages/daily-activities/daily-activities-list.html", data)
	}
}
