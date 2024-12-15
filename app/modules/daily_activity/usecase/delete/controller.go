package delete

import (
	"ipr/modules/daily_activity/authorization"
	"ipr/modules/daily_activity/repository"
	"net/http"
)

func Controller(repo *repository.DailyActivityRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		err := authorization.ErrorOnNotAuthorized(w, r, &id)
		if err != nil {
			return
		}

		err = repo.Delete(id)
		if err != nil {
			http.Error(w, "Failed to delete entity", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
