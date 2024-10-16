package service

import (
	"bufio"
	"log"
	"os"
	"strings"
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
	log.Println("Select menu", "\n 1 - Register", "\n 2 - Login")
	selected := Prompt("")

	switch selected {
	case "1":
		Name := Prompt("Your name :")
		Password := Prompt("Password :")

	default:
		return
	}
}
