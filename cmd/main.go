package main

import (
	_ "gin-boilerplate/docs"
	"gin-boilerplate/server"
)

// @title Gin Boilerplate
// @version v1.0.0
// @description This is a Boilerplate.
// @termsOfService http://example.com/terms/

// @contact.name API Support
// @contact.url http://www.example.com/support
// @contact.email support@example.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 127.0.0.1:8080
// @BasePath /api

func main() {
	// Start the server
	server.Run()
}
