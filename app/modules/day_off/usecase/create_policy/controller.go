package create_policy

import (
	"encoding/json"
	"ipr/infra/router/middleware"
	"ipr/modules/day_off/authorization"
	"ipr/shared"
	"net/http"
)

// Controller
// @Summary Create a new day off policy
// @Description This endpoint allows creating a new resource with a provided JSON payload.
// @Tags day-off
// @Accept json
// @Produce json
// @Param request body Command true "Request payload"
// @Success 201 {object} map[string]string "Created resource Id"
// @Failure 400 {object} map[string]string "Invalid JSON"
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /api/day-off/policies [post]
func Controller(h *Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := authorization.ErrorOnPolicyNotAuthorized(w, r, nil)
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
			Id string `json:"id"`
		}{Id: createdID})
	}
}
