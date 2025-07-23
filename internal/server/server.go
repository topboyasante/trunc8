package server

import (
	"net/http"

	"github.com/topboyasante/trunc8/internal/config"
	"github.com/topboyasante/trunc8/internal/handlers"
)

func InitServer(cfg *config.Config) *http.Server {
	mux := http.NewServeMux()

	mux.HandleFunc("/shorten", handlers.ShortenURL)
	mux.HandleFunc("/{code}", handlers.RedirectURL)

	server := &http.Server{
		Addr:    ":" + cfg.Server.Port,
		Handler: mux,
	}
	return server
}
