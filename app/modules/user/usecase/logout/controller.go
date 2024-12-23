package logout

import (
	"ipr/infra/session"
	"net/http"
)

func Controller(sm *session.Manager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sm.ClearUser(w, r)

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
