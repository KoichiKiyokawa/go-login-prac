package repository

import (
	"errors"

	"github.com/go-login-prac/src/entity"
	"golang.org/x/crypto/bcrypt"
)

type IUserRepository interface {
	FindByEmail(id string) (entity.User, error)
}

type UserRepository struct{}

func NewAuthRepository() UserRepository {
	return UserRepository{}
}

var hogehogePassword, _ = bcrypt.GenerateFromPassword([]byte("hogehoge"), 10)

var users = []entity.User{
	{Name: "user1", Email: "hoge@example.com", Password: string(hogehogePassword)},
	{Name: "user2", Email: "hoge2@example.com", Password: string(hogehogePassword)},
}

func (UserRepository) FindByEmail(email string) (entity.User, error) {
	for _, user := range users {
		if user.Email == email {
			return user, nil
		}
	}
	return entity.User{}, errors.New("not found")
}
