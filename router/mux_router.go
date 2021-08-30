package router

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gussf/go-bookstore/controller"
	"github.com/gussf/go-bookstore/model"
)

type MuxRouter struct {
	r           *mux.Router
	bookService controller.Controller
}

func NewMuxRouter(bc controller.Controller) *MuxRouter {
	r := mux.NewRouter()
	muxR := &MuxRouter{r, bc}
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
	WriteResponse(w, "Welcome to the go-bookstore API", http.StatusOK)
}

func (m *MuxRouter) FindById(w http.ResponseWriter, r *http.Request) {
	id := bookIdFromUrl(r)
	book, err := m.bookService.Find(id)

	if err != nil {
		WriteResponse(w, err.Error(), http.StatusNotFound)
		return
	}

	WriteResponse(w, book, http.StatusOK)
}

func (m *MuxRouter) All(w http.ResponseWriter, r *http.Request) {
	books, err := m.bookService.ListAll()
	if err != nil {
		WriteResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	WriteResponse(w, books, http.StatusOK)
}

func (m *MuxRouter) Add(w http.ResponseWriter, r *http.Request) {
	var book model.Book

	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		WriteResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = m.bookService.Validate(&book)
	if err != nil {
		WriteResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = m.bookService.Add(&book)
	if err != nil {
		WriteResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	WriteResponse(w, book, http.StatusCreated)
}

func (m *MuxRouter) RemoveById(w http.ResponseWriter, r *http.Request) {
	id := bookIdFromUrl(r)

	err := m.bookService.Remove(id)
	if err != nil {
		WriteResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	WriteResponse(w, "", http.StatusNoContent)
}

//Returns the book id sent by the client, parsed by the router
func bookIdFromUrl(r *http.Request) string {
	idStr := mux.Vars(r)["id"]
	return idStr
}
