package index

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("index中间件开始执行了")
		t := time.Now()
		c.Set("request", "中间件")
		status := c.Writer.Status()
		fmt.Println("中间件执行完毕", status)
		t2 := time.Since(t)
		fmt.Println("time:", t2)
	}
}
