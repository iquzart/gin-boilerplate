// Package server provides the implementation of the HTTP server.
package server

import (
	"context"
	"fmt"
	"gin-boilerplate/routes"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"
)

// serverConfigs represents the server configuration options.
type serverConfigs struct {
	port             string
	gracefulShutdown bool
}

// Start initializes the server and starts listening for incoming connections.
func Run() {
	// Get server configs
	serverConfigs := getConfigs()

	// Initialize the router
	router := routes.InitRouter()

	// Initialize the server
	server := &http.Server{
		Addr:    serverConfigs.port,
		Handler: router,
	}

	// Start the server
	if serverConfigs.gracefulShutdown {
		startWithGracefulShutdown(server)
	} else {
		start(server)
	}
}

// getConfigs returns the server configs.
func getConfigs() *serverConfigs {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	gracefulShutdown, err := strconv.ParseBool(os.Getenv("ENABLE_GRACEFUL_SHUTDOWN"))
	if err != nil {
		gracefulShutdown = true
	}
	return &serverConfigs{
		port:             fmt.Sprintf(":%s", port),
		gracefulShutdown: gracefulShutdown,
	}
}

// starts the server with out graceful shutdown.
func start(server *http.Server) {
	log.Printf("Server started on port %s", server.Addr)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Error starting server: %s\n", err)
	}
}

// starts the server with graceful shutdown.
func startWithGracefulShutdown(server *http.Server) {
	log.Printf("Started the server on port %s with graceful shutdown ", server.Addr)

	idleConnsClosed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint
		log.Println("Shutting down server...")

		// Give the server 5 seconds to finish processing active requests
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		// Shut down the server gracefully
		if err := server.Shutdown(ctx); err != nil {
			log.Fatalf("Error shutting down server: %s\n", err)
		}

		close(idleConnsClosed)
	}()

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Error starting server: %s\n", err)
	}

	<-idleConnsClosed
	log.Println("Server stopped")
}
