package data

import (
	"prime/models"

	"github.com/google/uuid"
)

var Users = []models.User{
	{
		ID:       uuid.New().String(),
		Name:     "Default",
		Password: "blabla",
	},
}
