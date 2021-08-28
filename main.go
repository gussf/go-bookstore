package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/gussf/go-bookstore/handlers"
	"github.com/gussf/go-bookstore/repository"
)

var addr = "0.0.0.0:15000"

func main() {

	repo, err := repository.NewPostgresRepo()
	if err != nil {
		panic(err)
	}
	books := handlers.NewBookHandler(repo)

	r := mux.NewRouter()
	s := r.PathPrefix("/api/v1").Subrouter()

	getRouter := s.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/", books.Index)
	getRouter.HandleFunc("/book/{id:[0-9]+}", books.FindById)
	getRouter.HandleFunc("/books", books.All)

	postRouter := s.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/book", books.Add)

	deleteRouter := s.Methods(http.MethodDelete).Subrouter()
	deleteRouter.HandleFunc("/book/{id:[0-9]+}", books.RemoveById)

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
	err = srv.Shutdown(ctx)
	if err != nil {
		panic(err)
	}
}
