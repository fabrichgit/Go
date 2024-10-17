package service

import (
	"bufio"
	"log"
	"os"
	"prime/src"
	"strings"

	"github.com/google/uuid"
)

func Prompt(message string) string {
	log.Println(message)

	// read user prompt
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	// clear after break line
	return strings.TrimSpace(input)
}

func Menu() {
	log.Println("\n", "\n", "Select menu", "\n 1 - Register", "\n 2 - Login", "\n 3 - All users", "\n 4 - Quit")
	selected := Prompt("")

	switch selected {
	case "1":
		Name := Prompt("Your name :")
		Password := Prompt("Password :")

		user := src.TUser{
			Id:       uuid.New().String(),
			Name:     Name,
			Password: hashPassword(Password),
		}

		Register(user)
		log.Println("\n", "\n", "Success 🥳!", "\n", src.Users, "\n", "\n", "")
	case "2":
		Name := Prompt("Your name :")
		Password := Prompt("Password :")

		if Login(Name, Password) {
			token := GenerateJWT(Name)
			log.Println("Token", "<", token, ">", "\n", "\n", "")
			log.Println("\n", "\n", "Success 🥳!", "\n", "\n", "")
		} else {
			log.Println("Invalid !")
		}

	case "3":
		AllUsers()
	case "4":
		return
	default:
		log.Println("Invalid choice")
	}

	Menu()
}
