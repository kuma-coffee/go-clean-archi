package main

import (
	"database/sql"
	"fmt"

	route "github.com/kuma-coffee/go-clean-archi/delivery"
	"github.com/kuma-coffee/go-clean-archi/repository"
	"github.com/kuma-coffee/go-clean-archi/usecase"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	username = "postgres"
	password = "postgres"
	dbName   = "test"
	port     = 5432
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, username, password, dbName)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	e := echo.New()

	bookRepo := repository.NewBookRepo(db)
	bookUsecase := usecase.NewBookUsecase(bookRepo)
	bookHandler := route.NewBookHandler(bookUsecase)

	e.POST("/books", bookHandler.AddBook)
	e.GET("/books", bookHandler.GetAllBooks)

	e.Start(":8080")
}
