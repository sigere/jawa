package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	configureRoutes(r)
	_ = r.Run(":8080")
}
