package repository

import (
	"database/sql"

	"github.com/kuma-coffee/go-clean-archi/entities"
)

type BookRepository interface {
	Store(book *entities.Book) error
	Fetch() ([]entities.Book, error)
	Update(id int, book *entities.Book) error
	Delete(id int) error
}

type bookRepository struct {
	db *sql.DB
}

func NewBookRepo(db *sql.DB) *bookRepository {
	return &bookRepository{db}
}

func (b *bookRepository) Store(book *entities.Book) error {
	stmt := `insert into "books"("id", "name", "year", "book_photo")values($1, $2, $3, $4)`

	_, err := b.db.Exec(stmt, book.ID, book.Name, book.Year, book.Photo)
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

func (b *bookRepository) Update(id int, book *entities.Book) error {
	stmt := `update books set name=$1, year=$2 where id=$3`

	_, err := b.db.Exec(stmt, book.Name, book.Year, id)
	if err != nil {
		return err
	}

	return nil
}

func (b *bookRepository) Delete(id int) error {
	stmt := `delete from books where id=$1`

	_, err := b.db.Exec(stmt, id)
	if err != nil {
		return err
	}

	return nil
}
