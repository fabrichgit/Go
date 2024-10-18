package main

import (
	"io"
	"log"
	"net/http"
	"prime/routes"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	http.HandleFunc("/hello", func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello, world!\n")
	})
	http.HandleFunc("/user/", routes.Userhandler)

	log.Println("About to listen on PORT :2005")

	err := http.ListenAndServe(":2005", nil)
	log.Fatal(err)
}
