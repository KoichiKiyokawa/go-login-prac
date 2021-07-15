package services

import (
	"errors"

	"github.com/go-login-prac/src/entities"
	"github.com/go-login-prac/src/repositories"
	"golang.org/x/crypto/bcrypt"
)

type IAuthService interface {
	ValidateUser(email string, password string) (entities.User, error)
}

type AuthService struct {
	AuthRepository repositories.IAuthRepository
}

func (s AuthService) ValidateUser(email string, password string) (entities.User, error) {
	commonError := errors.New("email or password is wrong") // メアドが登録されていることを攻撃者に知らせないように、エラーは統一する

	user, err := s.AuthRepository.FindByEmail(email)
	if err != nil {
		return entities.User{}, commonError // 該当するメアドのユーザーがいない
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return entities.User{}, commonError // パスワードが間違っている
	}
	return user, nil
}
