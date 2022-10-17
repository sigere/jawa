package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func getHandler(context *gin.Context) {
	id := context.Param("id")
	connection := getConnection()
	var p Product
	connection.First(&p, id)
	context.JSON(http.StatusOK, p)
}
