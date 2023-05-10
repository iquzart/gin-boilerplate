package routes

import (
	controllers "gin-boilerplate/controllers/api"

	"github.com/gin-gonic/gin"
)

func addAPIRoutes(rg *gin.RouterGroup) {
	api := rg.Group("/api")
	api.GET("/version", controllers.APIVersion)
}
