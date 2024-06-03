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
	e.GET("/home/homepage.html", Middleware(), homePage)
	e.GET("/admin/home/home.html", Middleware(), home)
	e.GET("/admin/home/console", Middleware(), console)
	e.GET("/admin/user/info", Middleware(), userInfo)
	e.GET("/admin/salary/list", Middleware(), salaryList)
	e.GET("/admin/attendance/list", Middleware(), attendanceList)
	e.GET("/admin/performance/list", Middleware(), performanceList)
	e.GET("/admin/leave/list", Middleware(), leaveList)
	e.GET("/admin/login", Middleware(), login)
	e.GET("/admin/login_out", loginOut)
	e.GET("/admin/workorder/list", Middleware(), workOrderList)
	e.GET("/admin/workorder/listform", Middleware(), workOrderListForm)
	e.POST("/admin/api/update_userinfo", Middleware(), updateUserInfo)
	e.GET("/admin", Middleware(), index)
	//e.GET("/", Middleware(), index)
}
