package create

import (
	"encoding/json"
	"net/http"
)

func HandleController(handler *UserCreateHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := &createUserRequest{
			Password: r.FormValue("password"),
			Email:    r.FormValue("email"),
		}

		id, err := handler.execute(req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(struct {
			ID string `json:"id"`
		}{ID: id})
	}
}
