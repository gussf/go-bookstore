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
)

var addr = "0.0.0.0:15000"

func main() {

	r := mux.NewRouter()
	s := r.PathPrefix("/api/v1").Subrouter()
	s.HandleFunc("/", handlers.Index)
	s.HandleFunc("/book/{id:[0-9]+}", handlers.GetBook).Methods("GET")

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
			os.Exit(1)
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
