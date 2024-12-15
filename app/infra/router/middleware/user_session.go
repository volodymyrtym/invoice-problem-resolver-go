package middleware

import (
	"context"
	http2 "ipr/infra/session"
	"net/http"
)

func UserSessionIdMiddleware(sm *http2.SessionManager) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			userId, err := sm.GetUser(r)
			if err != nil || userId == "" {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			ctx := context.WithValue(r.Context(), "user", userId)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func GetUserIdFromRequest(r *http.Request) string {
	return r.Context().Value("user").(string)
}
