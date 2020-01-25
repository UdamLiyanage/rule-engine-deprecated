package main

import (
	"github.com/arangodb/go-driver"
	"github.com/gin-gonic/gin"
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
	println("Main Function")
}
