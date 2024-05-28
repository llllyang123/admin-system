package admin

import (
	"fmt"
	"net/http"
	"user-admin/utils"

	"github.com/gin-gonic/gin"
)

func helloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello www.topgoer.com",
	})
}

func index(c *gin.Context) {
	one := c.Query("one")
	two := c.Query("two")
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

func login(c *gin.Context) {
	fmt.Println("login: login/admin.html")
	c.HTML(http.StatusOK, "login/admin.html", gin.H{"title": "管理系统登陆页面"})
}

func loginOut(c *gin.Context) {
	utils.LoginOut(c)
	login(c)
}
