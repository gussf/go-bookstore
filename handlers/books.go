package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gussf/go-bookstore/database"
)

type Books struct {
	l *log.Logger
	c *database.Connection
}

func NewBooks(l *log.Logger, c *database.Connection) *Books {
	return &Books{l, c}
}

func (b *Books) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/books/")

	switch r.Method {
	case http.MethodGet:
		fmt.Println("GET Method detected")
		b.getBooks(rw, id)
	case http.MethodPost:
		fmt.Println("POST Method detected")
		rw.WriteHeader(http.StatusCreated)
		fmt.Println(r)
	default:
		rw.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (b *Books) getBooks(rw http.ResponseWriter, id string) {
	enc := json.NewEncoder(rw)

	if id != "" {
		books, err := b.c.SelectBookById(id)
		if err != nil {
			b.l.Println(err)
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		if books == nil {
			rw.WriteHeader(http.StatusNotFound)
			return
		}

		rw.WriteHeader(http.StatusOK)
		enc.Encode(books)
		return
	}

	books, err := b.c.SelectAllBooks()
	if err != nil {
		b.l.Println(err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(http.StatusOK)
	enc.Encode(books)
}
