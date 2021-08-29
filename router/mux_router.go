package router

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gussf/go-bookstore/model"
)

type MuxRouter struct {
	r    *mux.Router
	repo model.Repository
}

func NewMuxRouter(repo model.Repository) *MuxRouter {
	r := mux.NewRouter()
	muxR := &MuxRouter{r, repo}
	muxR.setRoutes()

	return muxR
}

func (m *MuxRouter) setRoutes() {
	s := m.r.PathPrefix("/api/v1").Subrouter()

	getRouter := s.Methods(http.MethodGet).Subrouter()
	postRouter := s.Methods(http.MethodPost).Subrouter()
	deleteRouter := s.Methods(http.MethodDelete).Subrouter()
	// putR := s.Methods(http.MethodPut).Subrouter()

	// GET Method
	getRouter.HandleFunc("/", m.Index)
	getRouter.HandleFunc("/book/{id:[0-9]+}", m.FindById)
	getRouter.HandleFunc("/books", m.All)

	// POST Method
	postRouter.HandleFunc("/book", m.Add)

	// DELETE Method
	deleteRouter.HandleFunc("/book/{id:[0-9]+}", m.RemoveById)
}

func (m *MuxRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	m.r.ServeHTTP(w, r)
}

func (m *MuxRouter) Index(w http.ResponseWriter, r *http.Request) {
	WriteJsonToBody(w, "Welcome to the go-bookstore API")
}

func (m *MuxRouter) FindById(w http.ResponseWriter, r *http.Request) {
	id := bookIdFromUrl(r)
	book, err := m.repo.Select(id)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	WriteJsonToBody(w, book)
}

func (m *MuxRouter) All(w http.ResponseWriter, r *http.Request) {
	books, err := m.repo.SelectAll()

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	WriteJsonToBody(w, books)
}

func (m *MuxRouter) Add(w http.ResponseWriter, r *http.Request) {

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

	err = m.repo.Insert(&book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	WriteJsonToBody(w, book)
}

func (m *MuxRouter) RemoveById(w http.ResponseWriter, r *http.Request) {
	id := bookIdFromUrl(r)

	err := m.repo.Delete(id)
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
