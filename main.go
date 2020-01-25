package main

import (
	"github.com/arangodb/go-driver"
)

var database driver.Database

func init() {
	databaseConnect()
}

func main() {
	println("Main Function")
}
