package router

import (
	"encoding/json"
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

func WriteResponse(w http.ResponseWriter, v interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(v) //nolint
}
