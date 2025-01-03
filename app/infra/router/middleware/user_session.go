package middleware

import (
	"context"
	"ipr/infra/session"
	"net/http"
	"regexp"
)

func UserSessionIdMiddleware(sm *session.Manager) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			includeRegex := regexp.MustCompile(`^/daily-activities|^/day-off|^/api/daily-activities|^/api/day-off`)
			if !includeRegex.MatchString(r.URL.Path) {
				next.ServeHTTP(w, r)
				return
			}

			userId, err := sm.GetUser(r)
			if err != nil || userId == "" {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			// Додаємо userId до контексту
			ctx := context.WithValue(r.Context(), "user", userId)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func GetUserIdFromRequest(r *http.Request) string {
	return r.Context().Value("user").(string)
}
