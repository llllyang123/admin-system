package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"user-admin/app/admin/controller"
	"user-admin/utils"
)

func Index(c *gin.Context) {
	//one := c.Query("one")
	//two := c.Query("two")
	one, two := controller.Index(c)
	c.HTML(http.StatusOK, "admin/index.html", gin.H{"title": "adminIndex", "one": one, "two": two})
}

func Home(c *gin.Context) {
	one := c.Query("one")
	two := c.Query("two")
	c.HTML(http.StatusOK, "admin/home.html", gin.H{"title": "adminHome", "one": one, "two": two})
}

func HomePage(c *gin.Context) {
	//c.JSON(http.StatusOK, gin.H{
	//	"message": "Hello www.topgoer.com",
	//})
	c.HTML(http.StatusOK, "public/homepage.html", gin.H{"title": "adminHome"})

}

func Test(c *gin.Context) {
	one := c.Query("one")
	two := c.Query("two")
	c.HTML(http.StatusOK, "admin/test.html", gin.H{"title": "adminHome", "one": one, "two": two})
}

func WorkOrderList(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/work_order_list.html", gin.H{"title": "adminWorkOrder"})
}

func WorkOrderListForm(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/work_order_list_form.html", gin.H{"title": "adminWorkOrderListForm"})
}

func Console(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/console.html", gin.H{"title": "adminConsole"})
}

func UserInfo(c *gin.Context) {
	userinfo, err := controller.UserInfo()
	if err {
		fmt.Println(err)
	}
	c.HTML(http.StatusOK, "admin/user_info.html",
		gin.H{
			"title":    "adminUserInfo",
			"username": userinfo.UserName,
			"nickname": userinfo.NickName,
			"avatar":   userinfo.Avatar,
			"email":    userinfo.Email,
			"phone":    userinfo.Phone,
			"sex":      userinfo.Sex,
		})
}

func SalaryList(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/salary_list.html", gin.H{"title": "adminPersonalList"})
}

func SalaryListForm(c *gin.Context) {

	c.HTML(http.StatusOK, "admin/salary_list_form.html", gin.H{"title": "adminPersonalListForm"})
}

func PerformanceList(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/performance_list.html", gin.H{"title": "adminPersonalList"})
}

func LeaveList(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/leave_list.html", gin.H{"title": "adminPersonalList"})
}

func AttendanceList(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/attendance_list.html", gin.H{"title": "adminPersonalList"})
}

func Login(c *gin.Context) {
	fmt.Println("login: login/admin.html")
	c.HTML(http.StatusOK, "login/admin.html", gin.H{"title": "管理系统登陆页面"})
}

func LoginOut(c *gin.Context) {
	utils.LoginOut(c)
	Login(c)
}
