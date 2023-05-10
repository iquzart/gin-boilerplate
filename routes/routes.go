// Package routes contains the application's routes, which map user requests to the appropriate controller and method.
package routes

import (
	"github.com/gin-gonic/gin"
)

// InitRouter initializes the Gin router and sets up the middleware and routes.
func InitRouter() *gin.Engine {
	// Create a new Gin router instance.
	router := gin.New()

	// Use the Gin logger middleware to log HTTP requests and responses.
	router.Use(gin.Logger())

	// Use the Gin recovery middleware to recover from panics and return an error response instead of crashing.
	router.Use(gin.Recovery())

	// Add the application's routes to the specified router group.
	getRoutes(router.Group(""))

	// Return the initialized Gin router.
	return router
}

// getRoutes adds the application's routes to the specified router group.
func getRoutes(rg *gin.RouterGroup) {
	// Add the routes for the System.
	addSystemRoutes(rg)
	// Add the routes for API controllers.
	addAPIRoutes(rg)
}
