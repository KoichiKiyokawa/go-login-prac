package service

import (
	"testing"

	"github.com/go-login-prac/entity"
	"github.com/go-login-prac/util"
	"gorm.io/gorm"
)

var dummyUser = entity.User{Email: "test@example.com", Name: "user1", Password: "testtest", Model: &gorm.Model{ID: 1}}

type dummyUserRepository struct{}

func (dummyUserRepository) FindByEmail(email string) (entity.User, error) {
	var encryptedPassword, _ = util.HashString(dummyUser.Password)
	dummyUser.Password = string(encryptedPassword)
	return dummyUser, nil
}
func (dummyUserRepository) Find(id int) (entity.User, error) {
	return entity.User{}, nil
}
func (dummyUserRepository) All() ([]entity.User, error) {
	return []entity.User{}, nil
}
func (dummyUserRepository) Create(user entity.User) (entity.User, error) {
	return user, nil
}

func TestCreateUser(t *testing.T) {
	t.Run("will hash password before creating", func(t *testing.T) {
		userService := NewUserService(dummyUserRepository{})
		creatingUser := entity.User{Model: &gorm.Model{ID: 1}, Name: "1", Email: "t@example.com", Password: "testtest"}
		got, _ := userService.CreateUser(creatingUser)
		if creatingUser.Password == got.Password {
			t.Errorf("password is not hashed")
		}
	})
}
