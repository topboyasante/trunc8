package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/topboyasante/trunc8/internal/services"
	"github.com/topboyasante/trunc8/internal/types"
)

type ShortnerHandler struct {
	service *services.ShortnerService
}

func NewShortnerHandler(service *services.ShortnerService) *ShortnerHandler {
	return &ShortnerHandler{
		service: service,
	}
}

func (h *ShortnerHandler) ShortenURL(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		var payload types.ShortenRequest

		// read and decode the body into a struct
		err := json.NewDecoder(r.Body).Decode(&payload)
		if err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		// this expects some context
		res, err := h.service.ShortenURL(r.Context(), payload.URL)
		if err != nil {
			fmt.Print(err)
			http.Error(w, "Error shortening url", http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		jsonRes, err := json.Marshal(res)
		if err != nil {
			http.Error(w, "Error encoding response", http.StatusInternalServerError)
			return
		}

		w.Write(jsonRes)

		// Always close the body with defer r.Body.Close() to prevent resource leaks
		defer r.Body.Close()
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Not found"))
		log.Println("Received request with unsupported method on /shorten")
	}
}

func (h *ShortnerHandler) RedirectURL(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		shortCode := r.URL.Path[1:] // Remove leading "/" to get "someCODE"
		if shortCode == "" {
			http.Error(w, "No short code provided", http.StatusBadRequest)
			return
		}

		url, err := h.service.RedirectURL(r.Context(), shortCode)
		if err != nil {
			http.Error(w, "Unable to retrieve full url", http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, url, http.StatusMovedPermanently)
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Not found"))
		log.Println("Received request with unsupported method on /shorten")
	}
}
