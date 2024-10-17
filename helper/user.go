package helper

import "prime/models"

func FilterUsers(users []models.User, test func(User models.User) bool) []models.User {
	var filtered []models.User
	for _, user := range users {
		if test(user) {
			filtered = append(filtered, user)
		}
	}
	return filtered
}

func FindUser(users []models.User, test func(User models.User) bool) *models.User {
	for _, user := range users {
		if test(user) {
			return &user
		}
	}
	return nil
}
