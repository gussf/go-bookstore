package main

import (
	"github.com/gussf/go-bookstore/repository"
	"github.com/gussf/go-bookstore/router"
	"github.com/gussf/go-bookstore/server"
)

func main() {

	repo, err := repository.NewPostgresRepo()
	if err != nil {
		panic(err)
	}

	router := router.NewMuxRouter(repo)

	server := server.NewBookstore(router)
	server.Run()
}
