package router

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	bookstore "github.com/gussf/go-bookstore/src"
	"github.com/gussf/go-bookstore/src/books"
)

type MuxRouter struct {
	r           *mux.Router
	bookService books.Service
}

func NewMuxRouter(bs books.Service) *MuxRouter {
	r := mux.NewRouter()
	muxR := &MuxRouter{r, bs}
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

	bResp := responseFromService(book)
	WriteResponse(w, bResp, http.StatusOK)
}

func (m *MuxRouter) All(w http.ResponseWriter, r *http.Request) {
	books, err := m.bookService.ListAll()

	if err != nil {
		WriteResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var bResp []bookstore.BookResponse
	for _, v := range books {
		b := responseFromService(&v)
		bResp = append(bResp, b)
	}

	WriteResponse(w, bResp, http.StatusOK)
}

func (m *MuxRouter) Add(w http.ResponseWriter, r *http.Request) {
	var book bookstore.BookRequest

	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		WriteResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	b, err := m.bookService.NewBook(book.Title, book.Author, book.Copies, book.Price)
	if err != nil {
		WriteResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	bdto, err := m.bookService.Add(b)
	if err != nil {
		WriteResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	WriteResponse(w, responseFromService(bdto), http.StatusCreated)
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

func responseFromService(b *bookstore.BookDTO) bookstore.BookResponse {
	return bookstore.BookResponse{
		ID:           b.ID,
		Title:        b.Title,
		Author:       b.Author,
		Copies:       b.Copies,
		Price:        b.Price,
		CreationDate: b.CreationDate,
		Route:        "Mux",
	}
}
