package routers

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Option func(*gin.Engine)

var options []Option
var R = gin.New()

// 定义中间
func MiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		fmt.Println("总路由中间件开始执行了")
		// 设置变量到Context的key中，可以通过Get()取
		c.Set("request", "中间件")
		status := c.Writer.Status()
		fmt.Println("中间件执行完毕", status)
		t2 := time.Since(t)
		fmt.Println("time:", t2)
	}
}

// 注册app的路由配置
func Include(opts ...Option) {
	options = append(options, opts...)
}

// 初始化
func Init() *gin.Engine {
	//r := gin.New()
	// 自定义标识符，因为默认{{}}这种标识在前端框架中也有使用，会产生冲突
	// 注意：改完标识符之后别忘了把模板里原先的标识符一起改掉
	R.Delims("{[", "]}")
	R.Static("static/assets", "static/assets")
	R.LoadHTMLGlob("templates/**/*")
	// 加载404错误页面
	R.NoRoute(func(c *gin.Context) {
		// 实现内部重定向
		c.HTML(http.StatusOK, "error/404.html", gin.H{
			"title": "404",
		})
	})
	// add sessions
	store := cookie.NewStore([]byte("secret"))
	R.Use(sessions.Sessions("mysession", store))
	// 注册中间件
	R.Use(MiddleWare())
	for _, opt := range options {
		opt(R)
	}
	R.GET("/404", func(c *gin.Context) {
		c.HTML(http.StatusOK, "error/404.html", gin.H{"title": "404"})
	})
	return R
}
