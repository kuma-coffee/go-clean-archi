package main

import (
	"github.com/kuma-coffee/go-clean-archi/db"
	route "github.com/kuma-coffee/go-clean-archi/delivery"
	"github.com/kuma-coffee/go-clean-archi/repository"
	"github.com/kuma-coffee/go-clean-archi/usecase"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)

func main() {

	db := db.CreateConn()

	e := echo.New()

	bookRepo := repository.NewBookRepo(db)
	bookUsecase := usecase.NewBookUsecase(bookRepo)
	bookHandler := route.NewBookHandler(bookUsecase)

	e.POST("/books", bookHandler.AddBook)
	e.GET("/books", bookHandler.GetAllBooks)
	e.PUT("/books/:id", bookHandler.UpdateBook)
	e.DELETE("/books/:id", bookHandler.DeleteBook)

	e.Start(":8080")
}
