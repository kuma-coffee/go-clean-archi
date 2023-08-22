package repository

import (
	"database/sql"

	"github.com/kuma-coffee/go-clean-archi/entities"
	"github.com/kuma-coffee/go-clean-archi/helpers"
)

type UserRepository interface {
	CheckLogin(username, password string) (bool, error)
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *userRepository {
	return &userRepository{db}
}

func (u *userRepository) CheckLogin(username, password string) (bool, error) {
	newUser := entities.User{}

	stmt := "SELECT * FROM users WHERE username = $1"

	err := u.db.QueryRow(stmt, username).Scan(&newUser.ID, &newUser.Username, &newUser.Password)
	if err != nil {
		return false, err
	}

	match, err := helpers.CheckPasswordHash(password, newUser.Password)
	if !match {
		return false, err
	}

	return true, nil
}
