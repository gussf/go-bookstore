package server

import (
	"log"
	"net/http"

	"github.com/gussf/go-bookstore/router"
)

type Bookstore struct {
	router router.Router
}

func NewBookstore(router router.Router) Bookstore {
	return Bookstore{router}
}

func (b *Bookstore) Run(addr string) {
	log.Fatal((http.ListenAndServe(addr, b.router)))
}
