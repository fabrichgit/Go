package src

import (
	"log"
	"prime/src/service"
)

type User struct {
	Id       string
	Name     string
	Password *string
}

func Users() {
	// var users []User;
	// return
	name := service.Prompt("Your name")
	log.Println("Bonjour", name)
}
