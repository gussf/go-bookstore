package repository

import (
	"errors"

	"github.com/gussf/go-bookstore/model"
)

// Postgres repository
type InMemRepository struct {
	BookList map[string]model.Book
}

func NewInMemoryRepo() (*InMemRepository, error) {
	m := make(map[string]model.Book)
	return &InMemRepository{BookList: m}, nil
}

func (im InMemRepository) SelectAll() ([]model.Book, error) {

	var bks []model.Book

	for _, b := range im.BookList {
		bks = append(bks, b)
	}

	return bks, nil
}

func (im InMemRepository) Select(id string) (*model.Book, error) {
	book, ok := im.BookList[id]
	if !ok {
		return nil, errors.New("book with id=" + id + " not found")
	}
	return &book, nil
}

func (im InMemRepository) Insert(b *model.Book) error {
	return nil
}

func (im InMemRepository) Delete(id string) error {
	return nil
}
