package controller

import (
	"encoding/json"
	"net/http"
	"prime/data"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(data.Users)
}
