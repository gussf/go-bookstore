package books

import (
	"fmt"

	bookstore "github.com/gussf/go-bookstore/src"
)

// This Service looks unnecessary but we should have functions here which would contain business logic that dont belong in repository

type Service struct {
	repo bookstore.Repository
}

func NewService(repo bookstore.Repository) Service {
	return Service{repo: repo}
}

func (bc Service) Add(book *bookstore.BookDTO) error {
	fmt.Println(book.Title)
	return bc.repo.Insert(book)
}

func (bc Service) Find(id string) (*bookstore.BookDTO, error) {
	return bc.repo.Select(id)
}

func (bc Service) Remove(id string) error {
	return bc.repo.Delete(id)
}

func (bc Service) ListAll() ([]bookstore.BookDTO, error) {
	return bc.repo.SelectAll()
}

// Could have more than just struct-validation which might not fit well inside "bookstore"
func (bc Service) Validate(b *bookstore.BookDTO) error {
	return b.Validate()
}
