package model

import (
	"time"

	"github.com/go-playground/validator"
)

var validate *validator.Validate

type Book struct {
	ID           int
	Title        string `json:"title" validate:"required"`
	Author       string `json:"author" validate:"required"`
	Copies       int    `json:"copies" validate:"required"`
	Price        int    `json:"price" validate:"required"`
	CreationDate time.Time
}

func NewBook(title string, author string, copies int, price int) (book *Book, err error) {
	book = &Book{Title: title, Author: author, Copies: copies, Price: price}

	err = book.Validate()
	if err != nil {
		return nil, err
	}
	return book, nil
}

func (b Book) Validate() error {
	if validate == nil {
		validate = validator.New()
	}

	return validate.Struct(b)
}
