package create

import (
	"encoding/json"
	"ipr/shared"
	"net/http"
)

func HandleController(handler *UserCreateHandler) http.HandlerFunc {
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

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(struct {
			Id string `json:"id"`
		}{Id: id})
	}
}
