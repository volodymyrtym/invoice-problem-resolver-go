package middleware

import "net/http"

type Middleware func(http.Handler) http.Handler

// Apply applies a chain of middleware to a http
func Apply(handler http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	wrapped := http.Handler(handler)
	for i := len(middlewares) - 1; i >= 0; i-- {
		wrapped = middlewares[i](wrapped)
	}
	return wrapped.ServeHTTP
}
