package books

import (
	"github.com/go-playground/validator"
)

var validate *validator.Validate

type Book struct {
	Title  string `validate:"required"`
	Author string `validate:"required"`
	Copies int    `validate:"required"`
	Price  int64  `validate:"required"`
}

func NewBook(title string, author string, copies int, price int64) (book *Book, err error) {
	book = &Book{Title: title, Author: author, Copies: copies, Price: price}

	err = book.Validate()
	if err != nil {
		return nil, ErrNewBookValidationFailed
	}
	return book, nil
}

func (b Book) Validate() error {
	if validate == nil {
		validate = validator.New()
	}

	return validate.Struct(b)
}
