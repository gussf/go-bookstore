package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/gussf/go-bookstore/database"
	"github.com/gussf/go-bookstore/model"
)

var addr = "0.0.0.0:15000"

type Env struct {
	books database.BookModel
}

func main() {

	r := mux.NewRouter()
	db, err := database.NewConnection()
	if err != nil {
		log.Fatal(err)
	}

	env := &Env{
		books: database.BookModel{Conn: db},
	}
	fmt.Println(env)

	s := r.PathPrefix("/api/v1").Subrouter()
	s.HandleFunc("/", env.Index)
	s.HandleFunc("/book/{id:[0-9]+}", env.singleBook).Methods("GET")
	s.HandleFunc("/books", env.allBooks).Methods("GET")
	s.HandleFunc("/book", env.addBook).Methods("POST")

	srv := http.Server{
		Addr:         addr,
		Handler:      s,
		ErrorLog:     log.Default(),
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	fmt.Println("Starting server on", addr, "...")

	go func() {
		err := srv.ListenAndServe()
		if err != nil {
			panic(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	// Block until a signal is received.
	sig := <-c
	fmt.Println("Signal:", sig)

	ctx, f := context.WithTimeout(context.Background(), 30*time.Second)
	f()
	srv.Shutdown(ctx)

}

func (env *Env) Index(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Welcome to the go-bookstore API")
}

func (env *Env) singleBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	book, err := env.books.Find(vars["id"])

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(book)
}

func (env *Env) allBooks(w http.ResponseWriter, r *http.Request) {
	books, err := env.books.All()

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(books)
}

func (env *Env) addBook(w http.ResponseWriter, r *http.Request) {

	var book model.Book
	json.NewDecoder(r.Body).Decode(&book)

	err := book.Validate()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	err = env.books.Insert(&book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(book)
}
