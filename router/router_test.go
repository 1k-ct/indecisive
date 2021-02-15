package router

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
)

func testConnectRouter() error {
	router := InitRouters()
	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		return fmt.Errorf("error: %s", err)
	}
	router.ServeHTTP(w, req)
	return nil
}

var err error

func TestRouterIntroduction(t *testing.T) {
	router := InitRouters()
	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, w.Body.String(), "{\"msg\":\"you can serect some choicies.\"}")
}
func TestSelectOne(t *testing.T) {
	router := InitRouters()
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	body := bytes.NewBufferString("{\"items\":[\"a\",\"b\",\"c\"]}")
	c.Request, _ = http.NewRequest("POST", "/v1/api/select", body)
	router.ServeHTTP(w, c.Request)

	// assert.Equal(t,w.Body.String(),)
	assert.Equal(t, w.Code, 200)
}
