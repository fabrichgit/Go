package routes

import (
	"net/http"
	"prime/controller"
	"strings"
)

func Userhandler(res http.ResponseWriter, req *http.Request) {
	path := strings.TrimPrefix(req.URL.Path, "/user")

	switch {
	case path == "/all" && req.Method == http.MethodGet:
		controller.GetAllUsers(res, req)
	case path == "/register" && req.Method == http.MethodPost:
		controller.Register(res, req)
	case path == "/login" && req.Method == http.MethodPost:
		controller.Login(res, req)
	default:
		http.Error(res, "Invalid request method", http.StatusMethodNotAllowed)
	}
}
