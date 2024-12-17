package login

import (
	"encoding/json"
	"ipr/infra/session"
	"ipr/infra/template"
	"ipr/shared"
	"net/http"
)

func RenderController() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		template.RenderTemplate(w, "pages/login/login.html", nil)
	}
}

func HandlerController(handler *UserLoginHandler, sm *session.Manager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var command command
		if err := json.NewDecoder(r.Body).Decode(&command); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		id, err := handler.execute(&command)
		if err != nil {
			shared.HandleHttpError(w, r, err, nil)
			return
		}

		sm.SetUser(w, r, id)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(struct {
			Id string
		}{Id: id})
	}
}
