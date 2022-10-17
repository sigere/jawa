package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func configureRoutes(router *gin.Engine) {
	router.POST("/create/", createHandler)
	router.GET("/get/:id", getHandler)
	router.GET("/", notFoundHandler)
}

func notFoundHandler(context *gin.Context) {
	context.JSON(http.StatusNotFound, map[string]string{
		"error": "Page was not found :(",
	})
}
