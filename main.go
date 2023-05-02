package main

import (
	"gin-boilerplate/server"
	"log"
	"os"
)

func main() {
	// Custom port
	port := ":" + os.Getenv("PORT")

	// Initialize the server
	server := server.InitServer()

	// Run server
	if port == ":" {
		log.Printf("Gin App started on default port")
		server.Run()
	} else {
		log.Printf("Gin App started on port " + os.Getenv("PORT"))
		server.Run(port)
	}

}
