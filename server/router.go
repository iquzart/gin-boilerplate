package server

import (
	controllers "gin-boilerplate/controllers"

	"github.com/gin-gonic/gin"
)

func InitializeRoutes(r *gin.Engine) {
	// Routes
	r.GET("/health", controllers.Health)
}
