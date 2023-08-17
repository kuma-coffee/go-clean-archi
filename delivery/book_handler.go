package route

import (
	"net/http"

	"github.com/kuma-coffee/go-clean-archi/entities"
	"github.com/kuma-coffee/go-clean-archi/usecase"
	"github.com/labstack/echo/v4"
)

type BookHandler interface {
	AddBook(c echo.Context) error
	GetAllBooks(c echo.Context) error
}

type bookHandler struct {
	bookHandler usecase.BookUsecase
}

func NewBookHandler(bookUsecase usecase.BookUsecase) *bookHandler {
	return &bookHandler{bookUsecase}
}

func (h *bookHandler) AddBook(c echo.Context) error {
	var newBook entities.Book

	err := c.Bind(&newBook)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	err = h.bookHandler.Store(&newBook)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, "success create book")
}
func (h *bookHandler) GetAllBooks(c echo.Context) error {
	books, err := h.bookHandler.Fetch()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, books)
}
