package main

import (
	"log"
	"net/http"

	"github.com/gussf/go-bookstore/src/books"
	"github.com/gussf/go-bookstore/src/repository"
	"github.com/gussf/go-bookstore/src/router"
)

var addr = "0.0.0.0:15000"

func main() {

	repo, err := repository.NewPostgresRepo()
	if err != nil {
		panic(err)
	}
	defer repo.CloseConnection()

	bookC := books.NewService(repo)

	router := router.NewMuxRouter(bookC)

	log.Fatal(http.ListenAndServe(addr, router))
}
