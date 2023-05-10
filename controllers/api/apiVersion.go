package controllers

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func APIVersion(ctx *gin.Context) {
	apiVersion := os.Getenv("API_VERSION")
	if apiVersion == "" {
		apiVersion = "v1.0.0"
	}

	ctx.String(
		http.StatusOK,
		apiVersion,
	)
}
