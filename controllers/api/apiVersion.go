// APIVersion is a controller method that returns the API version.
package controllers

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// @Summary API Version
// @Description Returns API Version.
// @Produce  plain
// @Success 200 {string} string "Current API version"
// @Router /api/version [get]
// @Tags App
func APIVersion(ctx *gin.Context) {
	// Get the API version from the "API_VERSION" environment variable, or use a default value of "v1.0.0" if not set.
	apiVersion := os.Getenv("API_VERSION")
	if apiVersion == "" {
		apiVersion = "v1.0.0"
	}

	// Respond with a 200 status code and the API version as the response body.
	ctx.String(
		http.StatusOK,
		apiVersion,
	)
}
