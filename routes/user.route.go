package routes

import (
	"net/http"
	"prime/controller"
	"prime/middleware"
	"strings"
)

func Userhandler(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/user")

	switch {
	case path == "/all" && r.Method == http.MethodGet:
		middleware.JwtGuard(controller.GetAllUsers).ServeHTTP(w, r)
	case path == "/":
		middleware.JwtGuard(controller.GetOne).ServeHTTP(w, r)
	case path == "/register" && r.Method == http.MethodPost:
		controller.Register(w, r)
	case path == "/login" && r.Method == http.MethodPost:
		controller.Login(w, r)
	default:
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}
