package main

import (
	"github.com/kuma-coffee/go-clean-archi/db"
	route "github.com/kuma-coffee/go-clean-archi/delivery/http"
	"github.com/kuma-coffee/go-clean-archi/middleware"
	"github.com/kuma-coffee/go-clean-archi/repository"
	"github.com/kuma-coffee/go-clean-archi/usecase"
	"github.com/labstack/echo"
	_ "github.com/lib/pq"
)

func main() {
	db := db.CreateConn()

	e := echo.New()

	e.Static("/", "public")

	bookRepo := repository.NewBookRepo(db)
	bookUsecase := usecase.NewBookUsecase(bookRepo)
	bookHandler := route.NewBookHandler(bookUsecase)

	e.POST("/books", bookHandler.AddBook)
	e.GET("/books", bookHandler.GetAllBooks, middleware.IsAuthenticated)
	e.PUT("/books/:id", bookHandler.UpdateBook)
	e.DELETE("/books/:id", bookHandler.DeleteBook)

	userRepo := repository.NewUserRepo(db)
	userUsecase := usecase.NewUserUsecase(userRepo)
	userHandler := route.NewUserHandler(userUsecase)

	e.GET("/generate-hash/:password", route.GenerateHashPassword)
	e.POST("/login", userHandler.CheckLogin)
	e.POST("/register", userHandler.Register)

	// uploadHandler := route.NewUploadHanlder()

	// e.Static("/", "public")
	// e.POST("/upload", uploadHandler.Upload)

	e.Logger.Fatal(e.Start(":8080"))
}
