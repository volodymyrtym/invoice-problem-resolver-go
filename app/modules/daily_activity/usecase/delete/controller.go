package delete

import (
	"ipr/modules/daily_activity/authorization"
	"ipr/modules/daily_activity/repository"
	"net/http"
)

// Controller
// @Summary	  Delete Daily Activity
// @Description  Deletes a daily activity by its ID.
// @Tags		 daily-activity
// @Param		id   query  string  true  "The ID of the daily activity"
// @Success	  204  "No Content"
// @Failure	  401  "Unauthorized"
// @Failure	  500  "Internal Server Error"
// @Router	   /daily-activity [delete]
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
