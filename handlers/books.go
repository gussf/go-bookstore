package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gussf/go-bookstore/database"
	"github.com/gussf/go-bookstore/model"
)

type Books struct {
	l *log.Logger
	c *database.Connection
}

func NewBooks(l *log.Logger, c *database.Connection) *Books {
	return &Books{l, c}
}

func (b *Books) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		fmt.Println("GET Method detected")
		id := strings.TrimPrefix(r.URL.Path, "/books/")
		b.getBooks(rw, id)
	case http.MethodPost:
		fmt.Println("POST Method detected")
		b.insertBooks(rw, r)
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

func (b *Books) insertBooks(rw http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var book model.Book

	err := decoder.Decode(&book)
	if err != nil {
		b.l.Println(err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	err = b.c.InsertBook(book)
	if err != nil {
		b.l.Println(err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	rw.WriteHeader(http.StatusCreated)
}
