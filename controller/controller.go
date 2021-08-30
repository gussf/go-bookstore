package controller

import (
	"github.com/gussf/go-bookstore/model"
)

type Controller interface {
	Add(b *model.Book) error
	Remove(id string) error
	Find(id string) (*model.Book, error)
	ListAll() ([]model.Book, error)
	Validate(b *model.Book) error
}
