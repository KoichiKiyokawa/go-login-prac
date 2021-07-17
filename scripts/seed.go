package main

import (
	"github.com/go-login-prac/entity"
	"github.com/go-login-prac/repository"
	"github.com/go-login-prac/util"
)

func hashString(pass string) string {
	hash, _ := util.HashString(pass)
	return hash
}

func main() {
	users := []entity.User{
		{Email: "hoge@example.com", Password: hashString("hogehoge")},
		{Email: "hoge2@example.com", Password: hashString("hogehoge2")},
		{Email: "hoge3@example.com", Password: hashString("hogehoge3")},
	}

	repository.SetupDB()
	db := repository.GetDB()

	if err := db.Create(users).Error; err != nil {
		panic(err)
	}
}
