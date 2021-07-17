package controller

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-login-prac/entity"
)

type authServiceMockAlwaysValidateTrue struct{}

func (authServiceMockAlwaysValidateTrue) ValidateUser(email string, password string) (entity.User, error) {
	return entity.User{}, nil
}

type authServiceMockAlwaysValidateFalse struct{}

func (authServiceMockAlwaysValidateFalse) ValidateUser(email string, password string) (entity.User, error) {
	return entity.User{}, errors.New("")
}

func TestAuthLogin(t *testing.T) {
	t.Run("write to session when pass validate", func(t *testing.T) {
		authController := NewAuthController(authServiceMockAlwaysValidateTrue{})
		body := bytes.NewBufferString("{\"email\":\"email\",\"password\":\"password\"}")
		req := httptest.NewRequest(http.MethodPost, "/auth/login", body)
		got := httptest.NewRecorder()
		authController.AuthLogin(got, req)

		if got.Code != http.StatusOK {
			t.Error()
		}
		session, _ := getSession(req)
		if session.Values[SESSION_USER_KEY] == nil {
			t.Errorf("session is empty in spite of login success")
		}
	})

	t.Run("do not write to session when not pass validate", func(t *testing.T) {
		authController := NewAuthController(authServiceMockAlwaysValidateFalse{})
		body := bytes.NewBufferString("{\"email\":\"email\",\"password\":\"password\"}")
		req := httptest.NewRequest(http.MethodPost, "/auth/login", body)
		got := httptest.NewRecorder()
		authController.AuthLogin(got, req)

		if got.Code == http.StatusOK {
			t.Error()
		}
		session, _ := getSession(req)
		if session.Values[SESSION_USER_KEY] != nil {
			t.Errorf("session is not empty in spite of login failed")
		}
	})

}
