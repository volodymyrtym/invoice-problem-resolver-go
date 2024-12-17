package create

import (
	"encoding/json"
	"ipr/infra/router/middleware"
	"ipr/modules/daily_activity/authorization"
	"ipr/shared"
	"net/http"
)

func Controller(h *Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := authorization.ErrorOnNotAuthorized(w, r, nil)
		if err != nil {
			return
		}

		var command Command
		if err := json.NewDecoder(r.Body).Decode(&command); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		command.UserID = middleware.GetUserIdFromRequest(r)
		createdID, err := h.execute(&command)
		if err != nil {
			shared.HandleHttpError(w, r, err, nil)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(struct {
			ID string `json:"id"`
		}{ID: createdID})
	}
}
