package login

import (
	"encoding/json"
	"ipr/infra/router/middleware"
	"ipr/infra/session"
	"ipr/infra/template"
	"net/http"
)

func RenderController() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := map[string]interface{}{
			"message": "Welcome to Invoice Problem Resolver",
		}
		template.RenderTemplate(w, "pages/login/login.html", data)
	}
}

func HandlerController(handler *UserLoginHandler, sm *session.SessionManager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := &command{
			Password: r.FormValue("password"),
			Email:    r.FormValue("email"),
		}

		id, err := handler.execute(req)
		if err != nil {
			r = middleware.AddErrorToContext(r, err)
			return
		}

		sm.SetUser(w, r, id)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(struct {
			ID string
		}{ID: id})
	}
}
