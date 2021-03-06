package service

import (
	"github.com/go-login-prac/entity"
	"github.com/go-login-prac/repository"
	"github.com/go-login-prac/util"
)

type IUserService interface {
	GetUser(id int) (entity.User, error)
	GetAllUsers() ([]entity.User, error)
	CreateUser(user entity.User) (entity.User, error)
}

type UserService struct {
	userRepository repository.IUserRepository
}

func NewUserService(userRepository repository.IUserRepository) UserService {
	return UserService{userRepository: userRepository}
}

func (s UserService) GetUser(id int) (entity.User, error) {
	user, err := s.userRepository.Find(id)
	if err != nil {
		return entity.User{}, err
	}

	return user, nil
}

func (s UserService) GetAllUsers() ([]entity.User, error) {
	users, err := s.userRepository.All()
	if err != nil {
		return []entity.User{}, err
	}
	return users, nil
}

func (s UserService) CreateUser(user entity.User) (entity.User, error) {
	encryptedPassword, err := util.HashString(user.Password)
	if err != nil {
		return entity.User{}, nil
	}
	user.Password = encryptedPassword
	createdUser, err := s.userRepository.Create(user)
	if err != nil {
		return entity.User{}, nil
	}
	return createdUser, nil
}
