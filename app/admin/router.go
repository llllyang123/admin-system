package admin

import (
	"github.com/gin-gonic/gin"
)

func Routers(e *gin.Engine) {
	e.GET("/post", Middleware(), helloHandler)
	e.GET("/comment", helloHandler)
	//e.GET("/admin", helloHandler)
	e.GET("/admin/test", test)
	e.GET("/admin/home", Middleware(), index)
	e.GET("/home/homepage.html", homePage)
	e.GET("/admin/home/home.html", home)
	e.GET("/admin/home/console", console)
	e.GET("/admin/login", Middleware(), login)
	e.GET("/admin/login_out", loginOut)
	e.GET("/admin/workorder/list", Middleware(), workOrderList)
	e.GET("/admin/workorder/listform", Middleware(), workOrderListForm)
	e.GET("/admin", Middleware(), index)
	//e.GET("/", Middleware(), index)
}
