package model

import "time"

type Book struct {
	ID           int
	Title        string
	Author       string
	Copies       int
	Price        int
	CreationDate time.Time
}
