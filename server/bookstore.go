package server

import (
	"log"
	"net/http"

	"github.com/gussf/go-bookstore/model"
)

var addr = "0.0.0.0:15000"

type Bookstore struct {
	router model.Router
}

func NewBookstore(router model.Router) Bookstore {
	return Bookstore{router}
}

func (b *Bookstore) Run() {
	log.Fatal((http.ListenAndServe(addr, b.router)))
}
