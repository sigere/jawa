package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func createHandler(context *gin.Context) {
	connection := getConnection()
	_ = connection.AutoMigrate(&Product{})
	p := &Product{
		Name:  "newProduct",
		Price: 0,
	}
	connection.Create(p)

	context.JSON(http.StatusOK, p)
}
