package routes

import (
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	getRoutes(router.Group(""))
	return router
}

func getRoutes(rg *gin.RouterGroup) {
	addSystemRoutes(rg)
	addAPIRoutes(rg)
}
