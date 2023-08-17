package usecase

import (
	"github.com/kuma-coffee/go-clean-archi/entities"
	"github.com/kuma-coffee/go-clean-archi/repository"
)

type BookUsecase interface {
	Store(book *entities.Book) error
	Fetch() ([]entities.Book, error)
}

type bookUsecase struct {
	bookRepository repository.BookRepository
}

func NewBookUsecase(bookRepository repository.BookRepository) *bookUsecase {
	return &bookUsecase{bookRepository}
}

func (u *bookUsecase) Store(book *entities.Book) error {
	return u.bookRepository.Store(book)
}
func (u *bookUsecase) Fetch() ([]entities.Book, error) {
	return u.bookRepository.Fetch()
}
