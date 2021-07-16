package repository

import (
	"errors"

	"github.com/go-login-prac/src/entity"
	"golang.org/x/crypto/bcrypt"
)

type IUserRepository interface {
	FindByEmail(email string) (entity.User, error)
	All() ([]entity.User, error)
	Create(user entity.User) (entity.User, error)
}

type UserRepository struct{}

func NewUserRepository() UserRepository {
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

func (UserRepository) All() ([]entity.User, error) {
	return users, nil
}

func (UserRepository) Create(user entity.User) (entity.User, error) {
	users = append(users, user)
	return user, nil
}
