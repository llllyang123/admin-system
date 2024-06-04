package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"user-admin/app/admin/controller"
)

func UpdateUserInfo(c *gin.Context) {
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

func ApiSalaryList(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": []string{},
		"msg":  "msg",
	})
}
