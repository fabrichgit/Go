package main

import (
	"io"
	"log"
	"net/http"
	"prime/routes"
)

func main() {
	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello, world!\n")
	}

	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/user/", routes.Userhandler)

	log.Println("About to listen on 2005")

	err := http.ListenAndServe(":2005", nil)
	log.Fatal(err)
}
