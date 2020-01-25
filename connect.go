package main

import (
	"context"
	driver "github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
)

func databaseConnect() {
	conn, err := http.NewConnection(http.ConnectionConfig{
		Endpoints: []string{"http://13.57.187.149:8529"},
	})
	if err != nil {
		panic(err)
	}
	c, err := driver.NewClient(driver.ClientConfig{
		Connection:     conn,
		Authentication: driver.BasicAuthentication("udam", "Udam1998"),
	})
	if err != nil {
		panic(err)
	}

	database, err = c.Database(context.TODO(), "Devices")
	if err != nil {
		panic(err)
	}
}
