package server

import (
	"github.com/gin-gonic/gin"
	ginprometheus "github.com/mcuadros/go-gin-prometheus"
)

// Initialize server and routing information
func InitServer() *gin.Engine {

	// Creates a server without any middleware
	r := gin.New()

	// Add Logger and Recovery Middlewares
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	p := ginprometheus.NewPrometheus("gin")

	p.Use(r)

	// Initialize Routes
	InitializeRoutes(r)

	return r
}
