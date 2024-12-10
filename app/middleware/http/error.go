package http

import (
	"encoding/json"
	"html/template"
	"net/http"
)

// ErrorResponse defines a standard error response structure
type ErrorResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

// ErrorMiddleware handles errors and formats responses
func ErrorMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rec := recover(); rec != nil {
				respondWithError(w, r, http.StatusInternalServerError, "Internal Server Error")
			}
		}()

		next.ServeHTTP(w, r)
	})
}

// respondWithError sends a JSON or HTML error response
func respondWithError(w http.ResponseWriter, r *http.Request, status int, message string) {
	w.WriteHeader(status)

	// Check if the request accepts JSON
	if r.Header.Get("Accept") == "application/json" {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(ErrorResponse{
			Message: message,
			Code:    status,
		})
		return
	}

	// Render an HTML error template
	tmpl := template.Must(template.New("error").Parse(`
		<!DOCTYPE html>
		<html>
		<head><title>Error</title></head>
		<body>
			<h1>Error {{.Code}}</h1>
			<p>{{.Message}}</p>
		</body>
		</html>
	`))
	_ = tmpl.Execute(w, ErrorResponse{
		Message: message,
		Code:    status,
	})
}
