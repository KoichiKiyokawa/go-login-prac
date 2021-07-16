package repository

import (
	"errors"

	"github.com/go-login-prac/src/entity"
	"golang.org/x/crypto/bcrypt"
)

type IAuthRepository interface {
	FindByEmail(id string) (entity.User, error)
}

type AuthRepository struct{}

func NewAuthRepository() AuthRepository {
	return AuthRepository{}
}

var hogehogePassword, _ = bcrypt.GenerateFromPassword([]byte("hogehoge"), 10)

var users = []entity.User{
	{Name: "user1", Email: "hoge@example.com", Password: string(hogehogePassword)},
	{Name: "user2", Email: "hoge2@example.com", Password: string(hogehogePassword)},
}

func (AuthRepository) FindByEmail(email string) (entity.User, error) {
	for _, user := range users {
		if user.Email == email {
			return user, nil
		}
	}
	return entity.User{}, errors.New("not found")
}
