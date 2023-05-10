// Health is a controller method that returns a 200 status code with a "Working!" message
package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Health(ctx *gin.Context) {
	// Respond with a 200 status code and "Working!" message.
	ctx.String(
		http.StatusOK,
		"Working!")
}
