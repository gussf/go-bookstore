package controller

import (
	"github.com/gussf/go-bookstore/model"
	"github.com/gussf/go-bookstore/repository"
)

// This Service looks unnecessary but we should have functions here which would contain business logic that dont belong in repository

type bookService struct {
	repo repository.Repository
}

func NewBookService(repo repository.Repository) bookService {
	return bookService{repo: repo}
}

func (bc bookService) Add(book *model.Book) error {
	return bc.repo.Insert(book)
}

func (bc bookService) Find(id string) (*model.Book, error) {
	return bc.repo.Select(id)
}

func (bc bookService) Remove(id string) error {
	return bc.repo.Delete(id)
}

func (bc bookService) ListAll() ([]model.Book, error) {
	return bc.repo.SelectAll()
}

// Could have more than just struct-validation which might not fit well inside "model"
func (bc bookService) Validate(b *model.Book) error {
	return b.Validate()
}
