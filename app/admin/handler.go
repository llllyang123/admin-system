package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//已拆分至模块handler

func helloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello www.topgoer.com",
	})
}
