package bookstore

import "time"

type BookDTO struct {
	ID           int
	Title        string
	Author       string
	Copies       int
	Price        int
	CreationDate time.Time
}

type Repository interface {
	SelectAll() ([]BookDTO, error)
	Select(id string) (*BookDTO, error)
	Insert(*BookDTO) error
	Delete(id string) error
}
