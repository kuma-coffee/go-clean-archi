package route

import (
	"io"
	"mime/multipart"
	"os"

	"github.com/labstack/echo"
)

// type UploadHandler interface {
// 	Upload(c echo.Context) error
// 	OpenFile(fileName string, src multipart.File) error
// }

// type uploadHandler struct{}

// func NewUploadHanlder() *uploadHandler {
// 	return &uploadHandler{}
// }

func Upload(file multipart.FileHeader, c echo.Context) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// bs := make([]byte, file.Size)
	// _, err = bufio.NewReader(src).Read(bs)
	// if err != nil && err != io.EOF {
	// 	fmt.Println(err)
	// 	return nil, err
	// }

	dst, err := os.Create(file.Filename)
	if err != nil {
		return err
	}
	defer dst.Close()

	_, err = io.Copy(dst, src)
	if err != nil {
		return err
	}

	return nil
}

// func OpenFile(fileName string, src multipart.File) error {
// 	dst, err := os.Create(fileName)
// 	if err != nil {
// 		return err
// 	}
// 	defer dst.Close()

// 	_, err = io.Copy(dst, src)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
