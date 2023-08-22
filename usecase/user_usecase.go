package usecase

import (
	"github.com/kuma-coffee/go-clean-archi/repository"
)

type UserUsecase interface {
	CheckLogin(username, password string) (bool, error)
}

type userUsecase struct {
	userRepository repository.UserRepository
}

func NewUserUsecase(userRepository repository.UserRepository) *userUsecase {
	return &userUsecase{userRepository}
}

func (u *userUsecase) CheckLogin(username, password string) (bool, error) {
	return u.userRepository.CheckLogin(username, password)
}
