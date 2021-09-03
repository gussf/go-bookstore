package bookstore

import "time"

type BookDTO struct {
	ID           int
	Title        string
	Author       string
	Copies       int
	Price        int64
	CreationDate time.Time
}

type Repository interface {
	NewBook(title string, author string, copies int, price int64) *BookDTO
	SelectAll() ([]BookDTO, error)
	Select(id string) (*BookDTO, error)
	Insert(*BookDTO) error
	Delete(id string) error
	CloseConnection()
}
