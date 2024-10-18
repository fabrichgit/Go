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

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	// claims := r.Context().Value("user").(jwt.MapClaims)
	// if claims != nil {
	// 	userId := claims["id"]
	// 	json.NewEncoder(w).Encode(userId)
	// }
	json.NewEncoder(w).Encode(data.Users)
}

type AuthRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func Login(res http.ResponseWriter, req *http.Request) {
	body, err := io.ReadAll(req.Body)

	if err != nil {
		http.Error(res, "", http.StatusBadRequest)
		return
	}

	defer req.Body.Close()

	var loginReq AuthRequest
	err = json.Unmarshal(body, &loginReq)

	if err != nil {
		http.Error(res, "", http.StatusBadRequest)
		return
	}

	userFound := helper.FindUser(data.Users, func(u models.User) bool {
		return u.Name == loginReq.Name
	})

	if userFound != nil && helper.CheckPasswordHash(loginReq.Password, userFound.Password) {
		token := struct {
			Token string
		}{
			Token: helper.GenerateJWT(userFound.ID),
		}
		json.NewEncoder(res).Encode(token)
		return
	}

	http.Error(res, "", http.StatusUnauthorized)
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
	user.Password = helper.HashPassword(user.Password)

	data.Users = append(data.Users, user)

	json.NewEncoder(res).Encode(user)
}
