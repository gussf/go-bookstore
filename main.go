package main

import (
	"github.com/gussf/go-bookstore/controller"
	"github.com/gussf/go-bookstore/repository"
	"github.com/gussf/go-bookstore/router"
	"github.com/gussf/go-bookstore/server"
)

var addr = "0.0.0.0:15000"

func main() {

	repo, err := repository.NewPostgresRepo()
	if err != nil {
		panic(err)
	}

	bookC := controller.NewBookService(repo)

	router := router.NewMuxRouter(bookC)

	server := server.NewBookstore(router)
	server.Run(addr)
}
