package model

import "time"

type Book struct {
	ID           int
	Title        string `json:"title" validate:"required"`
	Author       string `json:"author" validate:"required"`
	Copies       int    `json:"copies" validate:"required"`
	Price        int    `json:"price" validate:"required"`
	CreationDate time.Time
}
