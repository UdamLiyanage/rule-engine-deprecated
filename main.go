package main

import (
	"github.com/arangodb/go-driver"
	"github.com/gin-gonic/gin"
	"log"
)

var database driver.Database

func init() {
	databaseConnect()
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/mqtt")

	return r
}

func main() {
	r := setupRouter()
	log.Fatal(r.Run(":9000"))
}
