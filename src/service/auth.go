package service

import (
	"errors"

	"github.com/go-login-prac/src/entity"
	"github.com/go-login-prac/src/repository"
	"golang.org/x/crypto/bcrypt"
)

type IAuthService interface {
	ValidateUser(email string, password string) (entity.User, error)
}

type AuthService struct {
	userRepository repository.IUserRepository
}

func NewAuthService(authRepository repository.IUserRepository) AuthService {
	return AuthService{userRepository: authRepository}
}

func (s AuthService) ValidateUser(email string, password string) (entity.User, error) {
	commonError := errors.New("email or password is wrong") // メアドが登録されていることを攻撃者に知らせないように、エラーは統一する

	user, err := s.userRepository.FindByEmail(email)
	if err != nil {
		return entity.User{}, commonError // 該当するメアドのユーザーがいない
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return entity.User{}, commonError // パスワードが間違っている
	}
	return user, nil
}
