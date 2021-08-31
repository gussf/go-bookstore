package model

import (
	"net/http"
)

type Router interface {
	ServeHTTP(http.ResponseWriter, *http.Request)
	Index(w http.ResponseWriter, r *http.Request)
	FindById(w http.ResponseWriter, r *http.Request)
	All(w http.ResponseWriter, r *http.Request)
	Add(w http.ResponseWriter, r *http.Request)
	RemoveById(w http.ResponseWriter, r *http.Request)
}
