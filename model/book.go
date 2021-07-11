package model

import "time"

type Book struct {
	ID           int
	Title        string `json:"title"`
	Author       string `json:"author"`
	Copies       int    `json:"copies"`
	Price        int    `json:"price"`
	CreationDate time.Time
}
