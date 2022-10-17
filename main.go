package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sigere/jawa/routing"
)

func main() {
	//connection := db.GetConnection()
	//err := model.Migrate(&connection)
	//if err != nil {
	//	log.Fatal(err)
	//	return
	//}

	//gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	routing.ConfigureRoutes(r)
	_ = r.Run(":8080")
}
