package route

import (
	"net/http"
	"strconv"

	"github.com/kuma-coffee/go-clean-archi/entities"
	"github.com/kuma-coffee/go-clean-archi/usecase"
	"github.com/labstack/echo"
)

type BookHandler interface {
	AddBook(c echo.Context) error
	GetAllBooks(c echo.Context) error
	UpdateBook(c echo.Context) error
	DeleteBook(c echo.Context) error
}

type bookHandler struct {
	bookUsecase usecase.BookUsecase
}

func NewBookHandler(bookUsecase usecase.BookUsecase) *bookHandler {
	return &bookHandler{bookUsecase}
}

func (h *bookHandler) AddBook(c echo.Context) error {
	// var newBook entities.Book

	name := c.FormValue("name")
	year := c.FormValue("year")
	file, err := c.FormFile("file")
	// file, err := c.MultipartForm()
	if err != nil {
		return err
	}

	// err := c.Bind(&newBook)
	// if err != nil {
	// 	return c.JSON(http.StatusInternalServerError, err.Error())
	// }

	err = Upload(*file, c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	newBook := entities.Book{Name: name, Year: year, Photo: ""}

	err = h.bookUsecase.Store(&newBook)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, "success create book")
}

func (h *bookHandler) GetAllBooks(c echo.Context) error {
	// err := middleware.ReadCookie(c)
	// if err != nil {
	// 	return c.JSON(http.StatusInternalServerError, err.Error())
	// }

	books, err := h.bookUsecase.Fetch()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, books)
}

func (h *bookHandler) UpdateBook(c echo.Context) error {
	var newBook entities.Book

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	err = c.Bind(&newBook)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	err = h.bookUsecase.Update(id, &newBook)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, "success update book")
}

func (h *bookHandler) DeleteBook(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	err = h.bookUsecase.Delete(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, "success delete book")
}
