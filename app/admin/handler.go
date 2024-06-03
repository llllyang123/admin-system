package admin

import (
	"fmt"
	"net/http"
	"user-admin/app/admin/controller"
	"user-admin/utils"

	"github.com/gin-gonic/gin"
)

func helloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello www.topgoer.com",
	})
}

func index(c *gin.Context) {
	//one := c.Query("one")
	//two := c.Query("two")
	one, two := controller.Index(c)
	c.HTML(http.StatusOK, "admin/index.html", gin.H{"title": "adminIndex", "one": one, "two": two})
}

func home(c *gin.Context) {
	one := c.Query("one")
	two := c.Query("two")
	c.HTML(http.StatusOK, "admin/home.html", gin.H{"title": "adminHome", "one": one, "two": two})
}

func homePage(c *gin.Context) {
	//c.JSON(http.StatusOK, gin.H{
	//	"message": "Hello www.topgoer.com",
	//})
	c.HTML(http.StatusOK, "public/homepage.html", gin.H{"title": "adminHome"})

}

func test(c *gin.Context) {
	one := c.Query("one")
	two := c.Query("two")
	c.HTML(http.StatusOK, "admin/test.html", gin.H{"title": "adminHome", "one": one, "two": two})
}

func workOrderList(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/work_order_list.html", gin.H{"title": "adminWorkOrder"})
}

func workOrderListForm(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/work_order_list_form.html", gin.H{"title": "adminWorkOrderListForm"})
}

func console(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/console.html", gin.H{"title": "adminConsole"})
}

func userInfo(c *gin.Context) {
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

func updateUserInfo(c *gin.Context) {
	_, err := controller.UpdateUserInfo(c)
	code := 0
	msg := "success"
	if err != "" {
		fmt.Println(err)
		code = 10000
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"data": []string{},
		"msg":  msg,
	})
}

func salaryList(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/salary_list.html", gin.H{"title": "adminPersonalList"})
}

func performanceList(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/performance_list.html", gin.H{"title": "adminPersonalList"})
}

func leaveList(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/leave_list.html", gin.H{"title": "adminPersonalList"})
}

func attendanceList(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/attendance_list.html", gin.H{"title": "adminPersonalList"})
}

func login(c *gin.Context) {
	fmt.Println("login: login/admin.html")
	c.HTML(http.StatusOK, "login/admin.html", gin.H{"title": "管理系统登陆页面"})
}

func loginOut(c *gin.Context) {
	utils.LoginOut(c)
	login(c)
}
