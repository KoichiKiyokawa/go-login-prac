package repository

import (
	"errors"

	"github.com/go-login-prac/src/model"
)

type IAuthRepository interface {
	FindByEmail(id string) (model.User, error)
}

type AuthRepository struct {
}

var users = []model.User{
	{Name: "user1", Email: "hoge@example.com", Password: "hogehoge"},
	{Name: "user2", Email: "hoge2@example.com", Password: "hogehoge"},
}

func (AuthRepository) FindByEmail(email string) (model.User, error) {
	for _, user := range users {
		if user.Email == email {
			return user, nil
		}
	}
	return model.User{}, errors.New("not found")
}
