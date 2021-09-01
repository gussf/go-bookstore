package bookstore

import (
	"net/http"
	"time"
)

type BookRequest struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Copies int    `json:"copies"`
	Price  int64  `json:"price"`
}

type BookResponse struct {
	ID           int       `json:"id"`
	Title        string    `json:"title"`
	Author       string    `json:"author"`
	Copies       int       `json:"copies"`
	Price        int64     `json:"price"`
	CreationDate time.Time `json:"created_at"`
	Route        string    `json:"route"`
}

type Router interface {
	ServeHTTP(http.ResponseWriter, *http.Request)
	Index(w http.ResponseWriter, r *http.Request)
	FindById(w http.ResponseWriter, r *http.Request)
	All(w http.ResponseWriter, r *http.Request)
	Add(w http.ResponseWriter, r *http.Request)
	RemoveById(w http.ResponseWriter, r *http.Request)
}
