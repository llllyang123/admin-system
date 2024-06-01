package view

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
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
	login(t, router)
}

func login(t *testing.T, router *gin.Engine) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/admin/login", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}
