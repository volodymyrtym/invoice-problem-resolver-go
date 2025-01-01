package middleware

import (
	"errors"
	"ipr/shared"
	"log"
	"net/http"
)

func ErrorHandlingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rec := recover(); rec != nil {
				log.Println("Recovered panic:", rec)
				err := errors.New("Unexpected server error")
				status := http.StatusInternalServerError
				shared.HandleHttpError(w, r, err, &status)
			}
		}()

		next.ServeHTTP(w, r)
	})
}
