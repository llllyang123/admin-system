package api

import (
	"fmt"
	"net/http"
	"strings"
	"user-admin/database"
	"user-admin/utils"

	"github.com/gin-gonic/gin"
)

type GetUser struct {
	db database.DB
}

var DB, nil = database.NewGormDB("database.db")

func helloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello www.topgoer.com",
	})
}

func login(c *gin.Context) {
	type UserInfo struct {
		UserName string `json:"username" form:"username"`
		Password string `json:"password" form:"password"`
		Age      string `json:"age" form:"age"`
	}

	userInfo := &UserInfo{}
	var users []utils.User

	if err := c.ShouldBind(&userInfo); err == nil {
		userName := strings.TrimSpace(userInfo.UserName)
		user := DB.Where("user_name = ?", userName).First(&users)
		if user.Error != nil {
			// 处理错误
			fmt.Println(user.Error)
			return
		}
		for _, user := range users {
			fmt.Printf("ID: %d, Name: %s,", user.Id, user.UserName)
		}
		if utils.MD5(users[0].Password) == utils.MD5(userInfo.Password) {
			utils.SessionUp(users, c)
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"data": []string{},
				"msg":  "成功",
				"user": users,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code": 10000,
				"data": []string{},
				"msg":  "密码错误",
				"user": users,
			})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
}
