package repository

import "github.com/gussf/go-bookstore/model"

type Repository interface {
	SelectAll() ([]model.Book, error)
	Select(id string) (*model.Book, error)
	Insert(*model.Book) error
	Delete(id string) error
}
