package view_test

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/index", func(c *gin.Context) {
		c.String(200, "pong")
	})
	r.GET("/admin", func(c *gin.Context) {
		c.String(200, "pong")
	})
	r.GET("/admin/login", func(c *gin.Context) {
		c.String(200, "pong")
	})
	return r
}

func TestPingRouter(t *testing.T) {
	router := setupRouter()
	index(t, router)
}

func index(t *testing.T, router *gin.Engine) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/index", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}
