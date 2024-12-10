package common

import (
	"ipr/middleware"
	"net/http"
)

type Route struct {
	Path    string
	Method  string
	Handler http.HandlerFunc
}

type Router struct {
	routes      []Route
	middlewares []middleware.Middleware
}

func NewRouter(middlewares ...middleware.Middleware) *Router {
	return &Router{middlewares: middlewares}
}

func (r *Router) AddRoute(path string, method string, handler http.HandlerFunc) {
	r.routes = append(r.routes, Route{Path: path, Method: method, Handler: handler})
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	for _, route := range r.routes {
		if route.Path == req.URL.Path && route.Method == req.Method {
			handler := middleware.Apply(route.Handler, r.middlewares...)
			handler(w, req)
			return
		}
	}
	http.NotFound(w, req)
}
