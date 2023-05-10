package routes

import (
	controllers "gin-boilerplate/controllers/system"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func addSystemRoutes(rg *gin.RouterGroup) {
	system := rg.Group("/system")
	system.GET("/health", controllers.Health)
	system.GET("/metrics", gin.WrapH(promhttp.Handler()))
}
