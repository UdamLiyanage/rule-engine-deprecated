package main

import (
	"context"
	"github.com/arangodb/go-driver"
	"github.com/gin-gonic/gin"
)

func inferRules(c *gin.Context) {
	getRules("test_serial")
	c.Status(200)
}

func getRules(serial string) {
	var res []map[string]interface{}
	bindVars := map[string]interface{}{
		"serial": serial,
	}
	cursor, err := database.Query(context.TODO(), "FOR r in rules FILTER r.serial==@serial RETURN r", bindVars)
	if err != nil {
		panic(err)
	}
	for {
		var doc map[string]interface{}
		_, err := cursor.ReadDocument(context.TODO(), &doc)
		if driver.IsNoMoreDocuments(err) {
			break
		} else if err != nil {
			panic(err)
		}
		res = append(res, doc)
	}
}
