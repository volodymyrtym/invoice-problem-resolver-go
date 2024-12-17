package create

import (
	"encoding/json"
	"ipr/shared"
	"net/http"
)

type UserCreateResponse struct {
	Id string `json:"id"`
}

// HandleController
// @Summary Create a new user
// @Description Accepts JSON input and creates a new user, returning the user ID.
// @Tags User
// @Accept json
// @Produce json
// @Param command body create.command true "User creation payload"
// @Success 201 {object} map[string]string "User ID"
// @Failure 400 {string} string "Invalid JSON"
// @Failure 500 {string} string "Internal Server Error"
// @Router /users [post]
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
		json.NewEncoder(w).Encode(UserCreateResponse{Id: id})
	}
}
