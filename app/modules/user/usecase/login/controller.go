package login

import (
	"encoding/json"
	"ipr/common"
	"net/http"
)

func RenderController() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := map[string]interface{}{
			"message": "Welcome to Invoice Problem Resolver",
		}
		common.RenderTemplate(w, "pages/login/login.html", data)
	}
}

func HandlerController(handler *UserLoginHandler, sessionManager *common.SessionManager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := &createCommand{
			Password: r.FormValue("password"),
			Email:    r.FormValue("email"),
		}

		id, err := handler.execute(req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		sessionManager.SetUser(w, r, id)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(struct {
			ID string
		}{ID: id})
	}
}
