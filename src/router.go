package bookstore

import (
	"net/http"
)

type BookRequest struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Copies int    `json:"copies"`
	Price  int    `json:"price"`
}

type BookResponse struct {
	ID           int    `json:"id"`
	Title        string `json:"title"`
	Author       string `json:"author"`
	Copies       int    `json:"copies"`
	Price        int    `json:"price"`
	CreationDate int    `json:"created_at"`
}

type Router interface {
	ServeHTTP(http.ResponseWriter, *http.Request)
	Index(w http.ResponseWriter, r *http.Request)
	FindById(w http.ResponseWriter, r *http.Request)
	All(w http.ResponseWriter, r *http.Request)
	Add(w http.ResponseWriter, r *http.Request)
	RemoveById(w http.ResponseWriter, r *http.Request)
}
