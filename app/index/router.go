package index

import "github.com/gin-gonic/gin"

func Routers(e *gin.Engine) {
	e.GET("/index/login", Middleware(), login)
	e.GET("/login", Middleware(), login)
	e.GET("/index", Middleware(), index)
	e.GET("/", Middleware(), index)
}
