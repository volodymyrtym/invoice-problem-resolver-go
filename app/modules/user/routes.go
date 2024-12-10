package user

import (
	"ipr/common"
	"ipr/modules/user/usecase/create"
	"ipr/modules/user/usecase/login"
	"net/http"
)

func RegisterRoutes(deps *Dependencies, router *common.Router, sessionManager *common.SessionManager) {
	router.AddRoute("/", http.MethodGet, login.RenderController())
	router.AddRoute("/users/create", http.MethodPost, create.HandleController(deps.CreateHandler))
	router.AddRoute("/users/login", http.MethodGet, login.RenderController())
	router.AddRoute("/users/login", http.MethodPut, login.HandlerController(deps.LoginHandler, sessionManager))
}
