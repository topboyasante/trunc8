package handlers

import (
	"log"
	"net/http"
)

func ShortenURL(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		log.Println("ShortenURL called with POST method")
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