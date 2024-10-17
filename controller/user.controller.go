package controller

import (
	"encoding/json"
	"net/http"
	"prime/data"
)

func GetAllUsers(res http.ResponseWriter, req *http.Request) {
	json.NewEncoder(res).Encode(data.Users)
}

func GetOneUsers(res http.Response, req *http.Request) {

}
