package handlers

import (
	"encoding/json"
	"net/http"
)

func WriteJsonToBody(w http.ResponseWriter, v interface{}) {
	err := json.NewEncoder(w).Encode(v)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
