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
		data := map[string]interface{}{
			"message": "Registration is turned off",
		}

		template.RenderTemplate(w, r, "pages/login/login.html", data)
	}
}

// HandlerController
// @Summary User Login Handler
// @Description Handles user login by processing login commands, setting session, and returning the user Id.
// @Tags User
// @Accept json
// @Produce json
// @Param command body command true "User login command"
// @Success 200 {object} map[string]string "User Id"
// @Failure 400 {string} string "Invalid JSON"
// @Failure 500 {object} string "Internal Server Error"
// @Router /users/login [post]
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
