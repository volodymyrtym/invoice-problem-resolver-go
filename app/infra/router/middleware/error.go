package middleware

import (
	"encoding/json"
	"errors"
	"fmt"
	"ipr/shared"
	"net/http"
	"strings"
)

func ErrorHandlingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rec := recover(); rec != nil {
				var err error
				if recErr, ok := rec.(error); ok {
					err = recErr
				} else {
					err = fmt.Errorf("%v", rec)
				}
				writeError(w, r, err)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func writeError(w http.ResponseWriter, r *http.Request, err error) {
	var statusCode int

	msg := "Internal server error"

	var invalidInputErr *shared.InvalidInputError
	if errors.As(err, &invalidInputErr) {
		statusCode = http.StatusBadRequest
		msg = invalidInputErr.Error()
	}

	if isJSONRequest(r) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(statusCode)
		json.NewEncoder(w).Encode(map[string]string{
			"error": msg,
		})
	} else {
		w.WriteHeader(statusCode)
		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprintf(w, "Error: %s\n", msg)
	}
}

func isJSONRequest(r *http.Request) bool {
	return strings.Contains(r.Header.Get("Accept"), "application/json") ||
		strings.Contains(r.Header.Get("Content-Type"), "application/json")
}
