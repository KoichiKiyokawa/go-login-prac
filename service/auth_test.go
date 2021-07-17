package service

import (
	"testing"
)

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
