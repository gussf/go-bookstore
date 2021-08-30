package server

import (
	"log"
	"net/http"

	"github.com/gussf/go-bookstore/model"
)

type Bookstore struct {
	router model.Router
}

func NewBookstore(router model.Router) Bookstore {
	return Bookstore{router}
}

func (b *Bookstore) Run(addr string) {
	log.Fatal((http.ListenAndServe(addr, b.router)))
}
