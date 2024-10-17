package routes

import (
	"fmt"
	"net/http"
	"prime/controller"
	"strings"
)

func Userhandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	fmt.Println(query)
	path := strings.TrimPrefix(r.URL.Path, "/user")

	switch {
	case path == "/all" && r.Method == http.MethodGet:
		controller.GetAllUsers(w, r)
	default:
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}
