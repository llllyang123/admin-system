package admin

import (
	"github.com/gin-gonic/gin"
	"user-admin/app/admin/handler"
)

func Routers(e *gin.Engine) {
	// template
	e.GET("/post", Middleware(), helloHandler)
	e.GET("/comment", helloHandler)
	//e.GET("/admin", helloHandler)
	e.GET("/admin/test", handler.Test)
	e.GET("/admin/home", Middleware(), handler.Index)
	e.GET("/home/homepage.html", Middleware(), handler.HomePage)
	e.GET("/admin/home/home.html", Middleware(), handler.Home)
	e.GET("/admin/home/console", Middleware(), handler.Console)
	e.GET("/admin/user/info", Middleware(), handler.UserInfo)
	e.GET("/admin/salary/list", Middleware(), handler.SalaryList)
	e.GET("/admin/salary/list_form", Middleware(), handler.SalaryListForm)
	e.GET("/admin/attendance/list", Middleware(), handler.AttendanceList)
	e.GET("/admin/performance/list", Middleware(), handler.PerformanceList)
	e.GET("/admin/leave/list", Middleware(), handler.LeaveList)
	e.GET("/admin/login", Middleware(), handler.Login)
	e.GET("/admin/login_out", handler.LoginOut)
	e.GET("/admin/workorder/list", Middleware(), handler.WorkOrderList)
	e.GET("/admin/workorder/listform", Middleware(), handler.WorkOrderListForm)

	//e.GET("/", Middleware(), index)
	// api
	e.POST("/admin/api/update_userinfo", Middleware(), handler.UpdateUserInfo)
	e.POST("/admin/api/salary/list", Middleware(), handler.ApiSalaryList)

	e.GET("/admin", Middleware(), handler.Index)
}
