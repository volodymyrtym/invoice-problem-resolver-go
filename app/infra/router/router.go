package router

import (
	"fmt"
	middleware2 "ipr/infra/router/middleware"
	"ipr/infra/session"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

var router *Router

type Route struct {
	Path    string
	Method  string
	Handler http.HandlerFunc
}

type Router struct {
	routes []Route
}

func InitializeRouter() {
	router = &Router{routes: []Route{}}
}

func AddRoute(path string, method string, handler http.HandlerFunc) {
	router.routes = append(router.routes, Route{Path: path, Method: method, Handler: handler})
}

func GetChiRouter(sm *session.Manager) http.Handler {
	chiRouter := chi.NewRouter()

	// Add middleware
	chiRouter.Use(middleware.Logger)
	chiRouter.Use(middleware.Recoverer)
	chiRouter.Use(middleware2.ErrorHandlingMiddleware)
	chiRouter.Use(middleware2.UserSessionIdMiddleware(sm))

	// Register routes
	for _, route := range router.routes {
		switch route.Method {
		case http.MethodGet:
			chiRouter.Get(route.Path, route.Handler)
		case http.MethodPost:
			chiRouter.Post(route.Path, route.Handler)
		case http.MethodPut:
			chiRouter.Put(route.Path, route.Handler)
		case http.MethodDelete:
			chiRouter.Delete(route.Path, route.Handler)
		default:
			fmt.Printf("Unsupported method %s for path %s\n", route.Method, route.Path)
		}
	}

	return chiRouter
}
