package books

import (
	"fmt"

	"github.com/gussf/go-bookstore/model"
)

// This Service looks unnecessary but we should have functions here which would contain business logic that dont belong in repository

type Service struct {
	repo model.Repository
}

func NewService(repo model.Repository) Service {
	return Service{repo: repo}
}

func (bc Service) Add(book *model.Book) error {
	fmt.Println(book.Title)
	return bc.repo.Insert(book)
}

func (bc Service) Find(id string) (*model.Book, error) {
	return bc.repo.Select(id)
}

func (bc Service) Remove(id string) error {
	return bc.repo.Delete(id)
}

func (bc Service) ListAll() ([]model.Book, error) {
	return bc.repo.SelectAll()
}

// Could have more than just struct-validation which might not fit well inside "model"
func (bc Service) Validate(b *model.Book) error {
	return b.Validate()
}
