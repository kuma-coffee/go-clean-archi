package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var IsAuthenticated = middleware.JWTWithConfig(middleware.JWTConfig{
	SigningKey: []byte("secret"),
})

func WriteCookie(token string, c echo.Context) error {
	cookie := new(http.Cookie)
	cookie.Name = "auth"
	cookie.Value = token
	cookie.Expires = time.Now().Add(24 * time.Hour)
	c.SetCookie(cookie)

	return c.String(http.StatusOK, "write a cookie")
}

func ReadCookie(c echo.Context) error {
	cookie, err := c.Cookie("auth")
	if err != nil {
		return err
	}

	fmt.Println(cookie.Name)
	fmt.Println(cookie.Value)

	return c.String(http.StatusOK, "read a cookie")
}
