package server

import (
	"net/http"

	"github.com/topboyasante/trunc8/internal/config"
	"github.com/topboyasante/trunc8/internal/handlers"
	"github.com/topboyasante/trunc8/internal/repositories"
	"github.com/topboyasante/trunc8/internal/services"
)

func InitServer(cfg *config.Config) *http.Server {
	// Initialize repository
	repo := repositories.NewShortnerRepository()

	// Initialize service with repository
	service := services.NewShortnerService(repo)

	// Initialize handler with service
	handler := handlers.NewShortnerHandler(service)

	mux := http.NewServeMux()

	mux.HandleFunc("/shorten", handler.ShortenURL)
	mux.HandleFunc("/{code}", handler.RedirectURL)

	server := &http.Server{
		Addr:    ":" + cfg.Server.Port,
		Handler: mux,
	}
	return server
}
