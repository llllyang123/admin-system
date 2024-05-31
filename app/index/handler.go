package index

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func index(c *gin.Context) {
	c.HTML(http.StatusOK, "index/index.html", gin.H{"title": "Index"})
}

func login(c *gin.Context) {
	fmt.Println("login: login/admin.html")
	c.HTML(http.StatusOK, "login/admin.html", gin.H{"title": "管理系统登陆页面"})
}
