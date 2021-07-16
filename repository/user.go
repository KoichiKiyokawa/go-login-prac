package repository

import (
	"errors"

	"github.com/go-login-prac/entity"
	"golang.org/x/crypto/bcrypt"
)

type IUserRepository interface {
	Find(id int) (entity.User, error)
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
	{ID: 1, Name: "user1", Email: "hoge@example.com", Password: string(hogehogePassword)},
	{ID: 2, Name: "user2", Email: "hoge2@example.com", Password: string(hogehogePassword)},
}

func (UserRepository) Find(id int) (entity.User, error) {
	for _, user := range users {
		if user.ID == id {
			return user, nil
		}
	}
	return entity.User{}, errors.New("not found")
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
	user.ID = users[len(users)-1].ID + 1
	users = append(users, user)
	return user, nil
}
