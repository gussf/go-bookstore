package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gussf/go-bookstore/model"
)

type BookHandler struct {
	repo model.Repository
}

func NewBookHandler(r model.Repository) (bh *BookHandler) {
	return &BookHandler{
		repo: r,
	}
}

func (bh *BookHandler) Index(w http.ResponseWriter, r *http.Request) {
	WriteJsonToBody(w, "Welcome to the go-bookstore API")
}

func (bh *BookHandler) FindById(w http.ResponseWriter, r *http.Request) {
	id := bookIdFromUrl(r)
	book, err := bh.repo.Select(id)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	WriteJsonToBody(w, book)
}

func (bh *BookHandler) All(w http.ResponseWriter, r *http.Request) {
	books, err := bh.repo.SelectAll()

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	WriteJsonToBody(w, books)
}

func (bh *BookHandler) Add(w http.ResponseWriter, r *http.Request) {

	var book model.Book

	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		WriteJsonToBody(w, err.Error())
		return
	}

	err = book.Validate()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		WriteJsonToBody(w, err.Error())
		return
	}

	err = bh.repo.Insert(&book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	WriteJsonToBody(w, book)
}

func (bh *BookHandler) RemoveById(w http.ResponseWriter, r *http.Request) {
	id := bookIdFromUrl(r)

	err := bh.repo.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// Returns the book id sent by the client, parsed by the router
func bookIdFromUrl(r *http.Request) string {
	idStr := mux.Vars(r)["id"]
	return idStr
}
