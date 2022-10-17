package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sigere/jawa/db"
	"github.com/sigere/jawa/model"
	"net/http"
)

func GetHandler(context *gin.Context) {
	id := context.Param("id")
	connection := db.GetConnection()
	var p model.Product
	connection.First(&p, id)
	context.JSON(http.StatusOK, p)
}
