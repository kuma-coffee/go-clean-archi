package usecase

import (
	"github.com/kuma-coffee/go-clean-archi/entities"
	"github.com/kuma-coffee/go-clean-archi/helpers"
	"github.com/kuma-coffee/go-clean-archi/repository"
)

type UserUsecase interface {
	CheckLogin(username, password string) (bool, error)
	Register(user *entities.User) error
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

func (u *userUsecase) Register(user *entities.User) error {
	passwordHash, err := helpers.HashPassword(user.Password)
	if err != nil {
		return err
	}

	user.Password = passwordHash

	return u.userRepository.Register(user)
}
