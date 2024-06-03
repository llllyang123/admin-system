package admin

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"user-admin/utils"
)

func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := utils.Redis.Get(context.Background(), "user")
		fmt.Println("redis_user_before:", user)
		utils.Redis.Set(context.Background(), "user", "admin", 0)
		user = utils.Redis.Get(context.Background(), "user")
		fmt.Println("redis_user_new:", user)
		// 暂时关闭校验方便开发
		//return
		fmt.Println("admin中间件开始执行了")
		t := time.Now()
		c.Set("request", "中间件")
		status := c.Writer.Status()
		//session := sessions.Default(c)
		if utils.Session == nil {
			utils.Init(c)
		}
		session := utils.Session
		fmt.Println("session:", session)
		fmt.Println("session_userid:", session.Get("userId"))
		fmt.Println("c.Request.URL.Path:", c.Request.URL.Path)

		if session.Get("userId") == nil {
			if c.Request.URL.Path != "/admin/login" {
				fmt.Println("未登录，请登录")
				fmt.Println("before_url: ", c.Request.URL.Path)
				session.Set("before_url", c.Request.URL.Path)
				c.Redirect(http.StatusFound, "/admin/login")
			}
		} else {
			if session.Get("before_url") != nil {
				targetURL := session.Get("before_url").(string)
				c.Redirect(http.StatusFound, targetURL)
				session.Delete("before_url")
			}
		}
		fmt.Println("session:", session.Get("userId"))
		fmt.Println("中间件执行完毕", status)
		t2 := time.Since(t)
		fmt.Println("time:", t2)
	}
}
