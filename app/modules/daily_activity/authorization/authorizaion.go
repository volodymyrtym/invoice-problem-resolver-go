package authorization

import (
	"github.com/pkg/errors"
	"ipr/modules/daily_activity/repository"
	"net/http"
)

type Auth struct {
	repo *repository.DailyActivityRepository
}

func NewAuth(repo *repository.DailyActivityRepository) *Auth {
	return &Auth{
		repo: repo,
	}
}

func (a *Auth) ErrorOnNotAuthorized(w http.ResponseWriter, r *http.Request, entityID string) error {
	userId, ok := r.Context().Value("user").(string)
	if !ok || userId == "" {
		http.Error(w, "no user in session", http.StatusUnauthorized)

		return errors.New("not authorized")
	}

	result, err := a.repo.IsOwner(entityID, userId)
	if err != nil {
		return err
	}

	if !result {
		http.Error(w, "not authorized", http.StatusUnauthorized)

		return errors.New("not authorized")
	}

	return nil
}
