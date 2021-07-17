package repository

import (
	"errors"

	"github.com/go-login-prac/entity"
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

func (UserRepository) Find(id int) (entity.User, error) {
	var user entity.User
	db.First(&user, id)
	return user, errors.New("not found")
}

func (UserRepository) FindByEmail(email string) (user entity.User, err error) {
	if err = db.Where("email = ?", email).Find(&user).Error; err != nil {
		return
	}
	return
}

func (UserRepository) All() (users []entity.User, err error) {
	if err = db.Find(&users).Error; err != nil {
		return
	}
	return
}

func (UserRepository) Create(user entity.User) (createdUser entity.User, err error) {
	if err = db.Create(user).Error; err != nil {
		return
	}
	return user, nil
}
