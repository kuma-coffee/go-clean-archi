package route

import (
	"io"
	"net/http"
	"os"

	"github.com/labstack/echo"
)

type UploadHandler interface {
	Upload(c echo.Context) error
}

type uploadHandler struct{}

func NewUploadHanlder() *uploadHandler {
	return &uploadHandler{}
}

func (u *uploadHandler) Upload(c echo.Context) error {
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}

	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	dst, err := os.Create(file.Filename)
	if err != nil {
		return err
	}
	defer dst.Close()

	_, err = io.Copy(dst, src)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, "file uploaded")
}
