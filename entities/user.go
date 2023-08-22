package entities

import "github.com/golang-jwt/jwt/v5"

type User struct {
	ID       int    `json:"user_id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type JWTCustomClaims struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	jwt.RegisteredClaims
}
