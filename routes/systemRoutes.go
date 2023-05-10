// Package routes contains the system routes, which map user requests to the appropriate controller and method.
package routes

import (
	controllers "gin-boilerplate/controllers/system"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// addSystemRoutes adds the routes for the System controller to the specified router group.
func addSystemRoutes(rg *gin.RouterGroup) {
	// Create a new group of routes for the System controller under the specified router group.
	system := rg.Group("/system")

	// Add a GET route for the Health method of the System controller.
	system.GET("/health", controllers.Health)

	// Add a GET route for the Metrics endpoint using the Prometheus HTTP handler.
	system.GET("/metrics", gin.WrapH(promhttp.Handler()))
}
