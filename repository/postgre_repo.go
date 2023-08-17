package repository

import (
	"database/sql"

	"github.com/kuma-coffee/go-clean-archi/entities"
)

type BookRepository interface {
	Store(book *entities.Book) error
	Fetch() ([]entities.Book, error)
}

type bookRepository struct {
	db *sql.DB
}

func NewBookRepo(db *sql.DB) *bookRepository {
	return &bookRepository{db}
}

func (b *bookRepository) Store(book *entities.Book) error {
	stmt := `insert into "books"("id", "name", "year")values($1, $2, $3)`

	_, err := b.db.Exec(stmt, book.ID, book.Name, book.Year)
	if err != nil {
		return err
	}

	return nil
}

func (b *bookRepository) Fetch() ([]entities.Book, error) {
	books := []entities.Book{}

	stmt := `select * from "books"`

	rows, err := b.db.Query(stmt)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		book := entities.Book{}
		err := rows.Scan(&book.ID, &book.Name, &book.Year)
		if err != nil {
			panic(err)
		}

		books = append(books, book)
	}

	return books, nil
}
