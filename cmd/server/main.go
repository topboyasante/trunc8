package main

import (
	"log"

	"github.com/topboyasante/trunc8/internal/config"
	"github.com/topboyasante/trunc8/internal/server"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to start server - configuration error: %v", err)
	}

	srv := server.InitServer(config)
	log.Printf("Starting server on port: %s", config.Server.Port)

	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}

	//  Anything below this will never execute because ListenAndServe() blocks indefinitely (unless it fails)
}
