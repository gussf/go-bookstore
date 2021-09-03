package repository

import (
	"errors"

	bookstore "github.com/gussf/go-bookstore/src"
)

// Postgres repository
type InMemRepository struct {
	BookList map[string]bookstore.BookDTO
}

func NewInMemoryRepo() (*InMemRepository, error) {
	m := make(map[string]bookstore.BookDTO)
	return &InMemRepository{BookList: m}, nil
}

func (im InMemRepository) NewBook(title string, author string, copies int, price int64) *bookstore.BookDTO {
	return &bookstore.BookDTO{
		Title:  title,
		Author: author,
		Copies: copies,
		Price:  price,
	}
}

func (im InMemRepository) SelectAll() ([]bookstore.BookDTO, error) {

	var bks []bookstore.BookDTO

	for _, b := range im.BookList {
		bks = append(bks, b)
	}

	return bks, nil
}

func (im InMemRepository) Select(id string) (*bookstore.BookDTO, error) {
	book, ok := im.BookList[id]
	if !ok {
		return nil, errors.New("book with id=" + id + " not found")
	}
	return &book, nil
}

func (im InMemRepository) Insert(b *bookstore.BookDTO) error {
	return nil
}

func (im InMemRepository) Delete(id string) error {
	return nil
}

func (im InMemRepository) CloseConnection() {
}
