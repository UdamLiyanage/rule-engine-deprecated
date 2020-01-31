package main

import (
	"context"
	"crypto/tls"
	driver "github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
	"os"
)

func databaseConnect() {
	conn, err := http.NewConnection(http.ConnectionConfig{
		Endpoints: []string{os.Getenv("ARANGODB_URL")},
		TLSConfig: &tls.Config{InsecureSkipVerify: true},
	})
	if err != nil {
		panic(err)
	}
	c, err := driver.NewClient(driver.ClientConfig{
		Connection:     conn,
		Authentication: driver.BasicAuthentication(os.Getenv("ARANGODB_USERNAME"), os.Getenv("ARANGODB_PASSWORD")),
	})
	if err != nil {
		panic(err)
	}

	database, err = c.Database(context.TODO(), os.Getenv("ARANGODB_DATABASE"))
	if err != nil {
		panic(err)
	}
}
