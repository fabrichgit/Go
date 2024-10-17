package controller

import (
	"encoding/json"
	"io"
	"net/http"
	"prime/data"
	"prime/helper"
	"prime/models"

	"github.com/google/uuid"
)

func GetAllUsers(res http.ResponseWriter, req *http.Request) {
	json.NewEncoder(res).Encode(data.Users)
}

func GetOneUsers(res http.Response, req *http.Request) {}

type UserAuth struct {
	name     string
	password string
}

type TokenPayload struct {
	token string
}

func Login(res http.ResponseWriter, req *http.Request) {
	body, err := io.ReadAll(req.Body)

	if err != nil {
		http.Error(res, "", http.StatusBadRequest)
		return
	}

	defer req.Body.Close()

	var user UserAuth
	err = json.Unmarshal(body, &user)

	if err != nil {
		http.Error(res, "", http.StatusBadRequest)
		return
	}

	var userFound = helper.FindUser(data.Users, func(u models.User) bool {
		return u.Name == user.name
	})

	if userFound == nil {
		http.Error(res, "", http.StatusNotFound)
		return
	}

	if helper.CheckPasswordHash(user.password, userFound.Password) {
		token := TokenPayload{
			token: helper.GenerateJWT(userFound.ID),
		}
		json.NewEncoder(res).Encode(token)
		return
	}

	http.Error(res, "", http.StatusBadRequest)
}

func Register(res http.ResponseWriter, req *http.Request) {
	body, err := io.ReadAll(req.Body)

	if err != nil {
		http.Error(res, "", http.StatusBadRequest)
		return
	}

	defer req.Body.Close()

	var user models.User
	err = json.Unmarshal(body, &user)

	if err != nil {
		http.Error(res, "", http.StatusBadRequest)
		return
	}

	user.ID = uuid.New().String()

	data.Users = append(data.Users, user)

	json.NewEncoder(res).Encode(user)
}
