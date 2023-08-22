package route

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/kuma-coffee/go-clean-archi/entities"
	"github.com/kuma-coffee/go-clean-archi/helpers"
	"github.com/kuma-coffee/go-clean-archi/usecase"
	"github.com/labstack/echo"
)

type UserHandler interface {
	CheckLogin(c echo.Context) error
}

type userHandler struct {
	userUsecase usecase.UserUsecase
}

func NewUserHandler(userUsecase usecase.UserUsecase) *userHandler {
	return &userHandler{userUsecase}
}

func GenerateHashPassword(c echo.Context) error {
	password := c.Param("password")

	hash, _ := helpers.HashPassword(password)

	return c.JSON(http.StatusOK, hash)
}

func (h *userHandler) CheckLogin(c echo.Context) error {
	var newUser entities.User

	err := c.Bind(&newUser)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	res, err := h.userUsecase.CheckLogin(newUser.Username, newUser.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	if !res {
		return echo.ErrUnauthorized
	}

	claims := &entities.JWTCustomClaims{
		Name:  newUser.Username,
		Admin: true,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
	})
}
