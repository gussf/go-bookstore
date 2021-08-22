package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gussf/go-bookstore/database"
	"github.com/gussf/go-bookstore/model"
)

const (
	REQUEST_SUCCESS = 1
)

type Books struct {
	l        *log.Logger
	c        *database.Connection
	validate *validator.Validate
}

var enc *json.Encoder
var dec *json.Decoder

func NewBooks(l *log.Logger, c *database.Connection) *Books {
	return &Books{l, c, validator.New()}
}

func (b *Books) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	enc = json.NewEncoder(rw)
	dec = json.NewDecoder(r.Body)

	var message string
	var httpCode int

	switch r.Method {
	case http.MethodGet:
		fmt.Println("GET Method detected")
		id := strings.TrimPrefix(r.URL.Path, "/books/")
		message, httpCode = b.getBooks(rw, id)
	case http.MethodPost:
		fmt.Println("POST Method detected")
		message, httpCode = b.insertBooks(rw, r)
	default:
		rw.WriteHeader(http.StatusMethodNotAllowed)
	}

	fmt.Println(message, httpCode)
	if httpCode != REQUEST_SUCCESS {
		resp := &model.ErrorReport{Message: message, HttpCode: httpCode}
		rw.WriteHeader(resp.HttpCode)
		enc.Encode(resp)
	}
}

func (b *Books) getBooks(rw http.ResponseWriter, id string) (string, int) {
	if id != "" {
		books, err := b.c.SelectBookById(id)
		if err != nil {
			return err.Error(), http.StatusInternalServerError
		}

		if books == nil {
			return "Book with id=" + id + " not found", http.StatusNotFound
		}

		rw.WriteHeader(http.StatusOK)
		enc.Encode(books)
		return "", REQUEST_SUCCESS
	}

	books, err := b.c.SelectAllBooks()
	if err != nil {
		return err.Error(), http.StatusInternalServerError
	}

	rw.WriteHeader(http.StatusOK)
	enc.Encode(books)
	return "", REQUEST_SUCCESS
}

func (b *Books) insertBooks(rw http.ResponseWriter, r *http.Request) (string, int) {
	var book model.Book

	err := dec.Decode(&book)
	if err != nil {
		return err.Error(), http.StatusBadRequest
	}

	err = b.validate.Struct(book)
	if err != nil {
		return err.Error(), http.StatusBadRequest
	}

	err = b.c.InsertBook(&book)
	if err != nil {
		return err.Error(), http.StatusBadRequest
	}

	rw.WriteHeader(http.StatusCreated)
	enc.Encode(book)
	return "", REQUEST_SUCCESS
}
