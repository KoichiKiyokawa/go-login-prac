package service

import (
	"testing"

	"github.com/go-login-prac/src/entity"
	"golang.org/x/crypto/bcrypt"
)

type dummyUserRepository struct{}

var dummyUser = entity.User{ID: 1, Email: "test@example.com", Name: "user1", Password: "testtest"}

func (r dummyUserRepository) FindByEmail(email string) (entity.User, error) {
	var encryptedPassword, _ = bcrypt.GenerateFromPassword([]byte(dummyUser.Password), 10)
	dummyUser.Password = string(encryptedPassword)
	return dummyUser, nil
}
func (r dummyUserRepository) Find(id int) (entity.User, error) {
	return entity.User{}, nil
}
func (r dummyUserRepository) All() ([]entity.User, error) {
	return []entity.User{}, nil
}
func (r dummyUserRepository) Create(user entity.User) (entity.User, error) {
	return entity.User{}, nil
}

func TestValidateUser(t *testing.T) {
	t.Run("password encrypt is correct", func(t *testing.T) {
		userService := NewAuthService(dummyUserRepository{})
		got, _ := userService.ValidateUser(dummyUser.Email, dummyUser.Password)
		want := dummyUser
		if got != want {
			t.Error()
		}
	})
}