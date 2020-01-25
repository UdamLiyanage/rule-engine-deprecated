package main

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"testing"
)

/*func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	DB.Collection = connect()
	os.Exit(m.Run())
}

func newRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Ping Successful",
		})
	})
	r.GET("/device-definitions/:id", readDefinition)
	r.POST("/device-definitions", createDefinition)
	r.PUT("/device-definitions/:id", updateDefinition)
	r.DELETE("/device-definitions/:id", deleteDefinition)
	return r
}*/

func testRequestStatusCode(method string, url string, body []byte, code int, t *testing.T) *httptest.ResponseRecorder {
	gin.SetMode(gin.TestMode)
	r := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(method, url, bytes.NewBuffer(body))
	r.ServeHTTP(w, req)
	if w.Code != code {
		t.Errorf("Status should be %d, got %d", code, w.Code)
	}
	return w
}

/*func testRequestBody(w *httptest.ResponseRecorder, field string, condition int, t *testing.T) {
	var response map[string]string
	_ = json.Unmarshal(w.Body.Bytes(), &response)
	value, exists := response[field]
	if !exists {
		t.Errorf("Wrong Response Format")
	}
	count, _ := strconv.Atoi(value)
	if count != condition {
		t.Errorf("Operation Not Working Properly!")
	}
}*/
