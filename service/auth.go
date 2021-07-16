package service

import (
	"errors"

	"github.com/go-login-prac/entity"
	"github.com/go-login-prac/repository"
	"github.com/go-login-prac/util"
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
	if util.CompareHashAndString(user.Password, password) {
		return user, nil
	}
	return entity.User{}, commonError // パスワードが間違っている
}
