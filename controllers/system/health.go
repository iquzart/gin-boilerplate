// Health is a controller method that returns a 200 status code with a "Working!" message
package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Check server health
// @Description Returns a "Working!" message if the server is healthy.
// @Produce  plain
// @Success 200 {string} string	"Working!"
// @Router /system/health [get]
// @Tags System
func Health(ctx *gin.Context) {
	// Respond with a 200 status code and "Working!" message.
	ctx.String(
		http.StatusOK,
		"Working!")
}
