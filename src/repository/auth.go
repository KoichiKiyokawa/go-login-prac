package repository

import (
	"errors"

	"github.com/go-login-prac/src/model"
	"golang.org/x/crypto/bcrypt"
)

type IAuthRepository interface {
	FindByEmail(id string) (model.User, error)
}

type AuthRepository struct{}

var hogehogePassword, _ = bcrypt.GenerateFromPassword([]byte("hogehoge"), 10)

var users = []model.User{
	{Name: "user1", Email: "hoge@example.com", Password: string(hogehogePassword)},
	{Name: "user2", Email: "hoge2@example.com", Password: string(hogehogePassword)},
}

func (AuthRepository) FindByEmail(email string) (model.User, error) {
	for _, user := range users {
		if user.Email == email {
			return user, nil
		}
	}
	return model.User{}, errors.New("not found")
}
