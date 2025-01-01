package delete

import (
	"ipr/infra/router"
	"ipr/modules/daily_activity/authorization"
	"ipr/modules/daily_activity/repository"
	"ipr/shared"
	"net/http"
)

// Controller
// @Summary	  Delete Daily Activity
// @Description  Deletes a daily activity by its Id.
// @Tags		 daily-activity
// @Param		id   query  string  true  "The Id of the daily activity"
// @Success	  204  "No Content"
// @Failure	  401  "Unauthorized"
// @Failure	  500  "Internal Server Error"
// @Router	   /api/daily-activities/{id} [delete]
func Controller(repo *repository.DailyActivityRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := router.GetURLParam("id", r)
		err := authorization.ErrorOnNotAuthorized(w, r, &id)
		if err != nil {
			shared.HandleHttpError(w, r, err, nil)
			return
		}
		err = repo.Delete(id)
		if err != nil {
			shared.HandleHttpError(w, r, err, nil)

			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
