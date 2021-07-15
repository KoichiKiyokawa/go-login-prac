package repositories

import (
	"errors"

	"github.com/go-login-prac/src/entities"

	"golang.org/x/crypto/bcrypt"
)

type IAuthRepository interface {
	FindByEmail(id string) (entities.User, error)
}

type AuthRepository struct{}

var hogehogePassword, _ = bcrypt.GenerateFromPassword([]byte("hogehoge"), 10)

var users = []entities.User{
	{Name: "user1", Email: "hoge@example.com", Password: string(hogehogePassword)},
	{Name: "user2", Email: "hoge2@example.com", Password: string(hogehogePassword)},
}

func (AuthRepository) FindByEmail(email string) (entities.User, error) {
	for _, user := range users {
		if user.Email == email {
			return user, nil
		}
	}
	return entities.User{}, errors.New("not found")
}
