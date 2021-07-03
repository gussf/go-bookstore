package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gussf/bookstore/handlers"
)

var addr = ":15000"

func main() {

	l := log.New(os.Stdout, "bookstore ", log.LstdFlags)

	bh := handlers.NewBooks(l)

	sm := http.NewServeMux()
	sm.Handle("/books", bh)

	s := http.Server{
		Addr:         addr,
		Handler:      sm,
		ErrorLog:     l,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	fmt.Println("Starting server on", addr, "...")

	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatalf("Failed to start server: %s\n", err)
			os.Exit(1)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	// Block until a signal is received.
	sig := <-c
	l.Println("Signal:", sig)

	ctx, f := context.WithTimeout(context.Background(), 30*time.Second)
	f()
	s.Shutdown(ctx)

}
