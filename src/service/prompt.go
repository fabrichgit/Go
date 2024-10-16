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
