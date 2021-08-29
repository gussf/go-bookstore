package router

import (
	"encoding/json"
	"net/http"
)

func WriteJsonToBody(w http.ResponseWriter, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(v)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
