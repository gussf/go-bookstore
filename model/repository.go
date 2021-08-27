package model

type Repository interface {
	SelectAll() ([]Book, error)
	Select(id string) (*Book, error)
	Insert(*Book) error
	Delete(id string) error
}
