package service

import (
	"log"
	"prime/src"
)

func Register(user src.TUser) {
	src.Users = append(src.Users, user)
}

func Login(name string, password string) bool {
	UserFound := findUser(src.Users, func(u src.TUser) bool {
		return u.Name == name
	})

	return (UserFound != nil && checkPasswordHash(password, UserFound.Password))
}

func AllUsers() {
	log.Println(src.Users)
}

// Utils functions
func filterUsers(users []src.TUser, test func(User src.TUser) bool) []src.TUser {
	var filtered []src.TUser
	for _, user := range users {
		if test(user) {
			filtered = append(filtered, user)
		}
	}
	return filtered
}

// Fonction pour trouver un utilisateur (équivalent de .find())
func findUser(users []src.TUser, test func(User src.TUser) bool) *src.TUser {
	for _, user := range users {
		if test(user) {
			return &user // renvoie un pointeur vers l'utilisateur trouvé
		}
	}
	return nil // renvoie nil si aucun utilisateur trouvé
}
