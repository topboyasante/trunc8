package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/topboyasante/trunc8/internal/config"
	"github.com/topboyasante/trunc8/internal/database"
	"github.com/topboyasante/trunc8/internal/server"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to start server - configuration error: %v", err)
	}

	err = database.ConnectToMongo(config)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}

	// Handle graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		log.Println("Shutting down...")
		database.DisconnectMongo()
		os.Exit(0)
	}()

	srv := server.InitServer(config)
	log.Printf("Starting server on port: %s", config.Server.Port)

	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
