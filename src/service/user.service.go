package service

import (
	"log"
	"prime/src"
)

func Register(user src.TUser) {
	append(src.Users, user)
	log.Println(src.Users)
}
