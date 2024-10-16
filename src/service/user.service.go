package service

import (
	"log"
	"prime/src"
)

func Register(user src.TUser) {
	src.Users = append(src.Users, user)
	log.Println(src.Users)
}

func AllUsers() {
	log.Println(src.Users)
}
