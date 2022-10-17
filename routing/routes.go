package routing

import (
	"github.com/gin-gonic/gin"
	"github.com/sigere/jawa/controller"
	"net/http"
)

func ConfigureRoutes(router *gin.Engine) {
	router.Group("/product")
	{
		router.POST("/create", controller.ProductCreateHandler)
		router.GET("/get/:id", controller.ProductGetHandler)
		router.GET("/update/:id", controller.ProductUpdateHandler)
		router.GET("/delete/:id", controller.ProductDeleteHandler)
	}

	router.GET("/", notFoundHandler)
}

func notFoundHandler(context *gin.Context) {
	context.JSON(http.StatusNotFound, map[string]string{
		"error": "Page was not found :(",
	})
}
