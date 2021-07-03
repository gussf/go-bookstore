package handlers

import (
	"fmt"
	"log"
	"net/http"
)

type Books struct {
	l *log.Logger
}

func NewBooks(l *log.Logger) *Books {
	return &Books{l}
}

func (b *Books) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case http.MethodGet:
		fmt.Println("GET Method detected")
		rw.WriteHeader(http.StatusOK)
		fmt.Println(r)
	case http.MethodPost:
		fmt.Println("POST Method detected")
		rw.WriteHeader(http.StatusCreated)
		fmt.Println(r)
	default:
		rw.WriteHeader(http.StatusMethodNotAllowed)
	}
}
