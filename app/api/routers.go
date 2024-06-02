package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

// 定义中间
func MiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		fmt.Println("api中间件开始执行了")
		// 设置变量到Context的key中，可以通过Get()取
		c.Set("request", "中间件")
		status := c.Writer.Status()
		fmt.Println("中间件执行完毕", status)
		t2 := time.Since(t)
		fmt.Println("time:", t2)
	}
}

// Routers
func Routers(e *gin.Engine) {
	e.GET("/api/login", MiddleWare(), login)
	e.GET("/checkout", helloHandler)
	e.GET("/api/test", test)
	e.POST("/api/edit_user_info", editUserInfo)
	e.GET("/api", MiddleWare(), helloHandler)
}
