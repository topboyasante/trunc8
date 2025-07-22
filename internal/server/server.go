package server

import (
	"log"
	"net/http"

	"github.com/topboyasante/trunc8/internal/config"
)

func InitServer(cfg *config.Config) *http.Server {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, Trunc8!"))
		log.Println("Received request on /")
	})

	server := &http.Server{
		Addr:    ":" + cfg.Server.Port,
		Handler: mux,
	}
	return server
}
