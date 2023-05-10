// Package routes contains the application routes, which map user requests to the appropriate controller and method.
package routes

import (
	controllers "gin-boilerplate/controllers/api"

	"github.com/gin-gonic/gin"
)

// addAPIRoutes adds the routes for the API controller to the specified router group.
func addAPIRoutes(rg *gin.RouterGroup) {
	// Create a new group of routes for the API controller under the specified router group.
	api := rg.Group("/api")

	// Add a GET route for the APIVersion method of the API controller.
	api.GET("/version", controllers.APIVersion)
}
