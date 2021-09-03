package books

import (
	bookstore "github.com/gussf/go-bookstore/src"
)

// This Service looks unnecessary but we should have functions here which would contain business logic that dont belong in repository

type Service struct {
	repo bookstore.Repository
}

func NewService(repo bookstore.Repository) Service {
	return Service{repo: repo}
}

func (bc Service) Add(b *Book) (*bookstore.BookDTO, error) {
	bDto := bc.repo.NewBook(b.Title, b.Author, b.Copies, b.Price)
	return bDto, bc.repo.Insert(bDto)
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

func (bc Service) NewBook(title string, author string, copies int, price int64) (book *Book, err error) {
	return NewBook(title, author, copies, price)
}
