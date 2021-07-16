package service

import (
	"testing"

	"github.com/go-login-prac/entity"
)

var dummyUser = entity.User{ID: 1, Email: "test@example.com", Name: "user1", Password: "testtest"}

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
