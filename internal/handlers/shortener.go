package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/topboyasante/trunc8/internal/types"
)

func ShortenURL(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		var payload types.ShortenRequest

		// read and decode the body into a struct
		err := json.NewDecoder(r.Body).Decode(&payload)
		if err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		fmt.Println(payload.URL)

		// Always close the body with defer r.Body.Close() to prevent resource leaks
		defer r.Body.Close()
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Not found"))
		log.Println("Received request with unsupported method on /shorten")
	}
}

func RedirectURL(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		log.Println("ShortenURL called with POST method")
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Not found"))
		log.Println("Received request with unsupported method on /shorten")
	}
}
