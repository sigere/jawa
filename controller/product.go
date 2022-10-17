package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sigere/jawa/db"
	"github.com/sigere/jawa/model"
	"gorm.io/gorm/clause"
	"log"
	"net/http"
)

func ProductCreateHandler(context *gin.Context) {
	connection := db.GetConnection()
	p := model.Product{}
	if err := context.ShouldBind(&p); err != nil {
		log.Println(err)
	}

	if errors := p.Validate(); len(errors) != 0 {
		context.JSON(http.StatusBadRequest, map[string]any{
			"success": false,
			"error":   errors,
		})
		return
	}

	connection.Omit("Manufacturer").Create(&p)
	context.JSON(http.StatusOK, map[string]any{
		"success":        true,
		"object created": p,
	})
}

func ProductGetHandler(context *gin.Context) {
	connection := db.GetConnection()
	id := context.Param("id")
	var p model.Product
	connection.Preload(clause.Associations).First(&p, id)

	context.JSON(http.StatusOK, p)
}

func ProductUpdateHandler(context *gin.Context) {
}

func ProductDeleteHandler(context *gin.Context) {
}
