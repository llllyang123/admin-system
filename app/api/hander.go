package api

import (
	"fmt"
	"net/http"
	"strings"
	apiController "user-admin/app/api/controller"
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

func test(c *gin.Context) {
	userInfo, count, _ := apiController.Test()
	fmt.Printf("json - userInfo: %d, count: %d", userInfo, count)
	c.JSON(http.StatusOK, gin.H{
		"code":  0,
		"data":  userInfo,
		"count": count,
		"msg":   "success",
	})
}

func getUserSalaryInfo(c *gin.Context) {
	userInfo, err := apiController.GetUserSalaryInfo(c)
	if !err {
		type SalaryInfo struct {
			JobId       uint
			Salary      float64
			BasicSalary float64
			NickName    string
		}
		var salaryInfo SalaryInfo
		salaryInfo.JobId = userInfo.JobId
		salaryInfo.Salary = userInfo.Salary
		salaryInfo.BasicSalary = userInfo.BasicSalary
		salaryInfo.NickName = userInfo.NickName
		c.JSON(http.StatusOK, gin.H{
			"code":  0,
			"data":  salaryInfo,
			"count": 1,
			"msg":   "success",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":  10000,
			"data":  []string{},
			"count": 0,
			"msg":   "error",
		})
	}
}

func editUserInfo(c *gin.Context) {
	status, err := apiController.EditUserInfo(c)
	code := 0
	msg := "success"
	if !status {
		code = 10000
		msg = "error"
	}
	if err != "" {
		msg = err
	}
	c.JSON(http.StatusOK, gin.H{
		"code":  code,
		"data":  []string{},
		"count": 1,
		"msg":   msg,
	})
}

func getSalaryList(c *gin.Context) {
	dataList, count, _, _ := apiController.GetSalaryList(c)
	c.JSON(http.StatusOK, gin.H{
		"code":  0,
		"data":  dataList,
		"count": count,
		"msg":   "success",
	})
}

func addSalary(c *gin.Context) {
	_, err, types := apiController.AddSalaryList(c)
	if err == nil {
		msg := "success"
		code := 0
		if types == 1 {
			code = 10000
			msg = "已对该员工所在月份发过工资或者该员工不存在"
		}
		c.JSON(http.StatusOK, gin.H{
			"code":  code,
			"data":  []string{},
			"count": 1,
			"msg":   msg,
		})
	}
}

func login(c *gin.Context) {
	type UserInfo struct {
		UserName string `json:"username" form:"username"`
		Password string `json:"password" form:"password"`
		Age      string `json:"age" form:"age"`
	}

	userInfo := &UserInfo{}
	var users utils.User

	if err := c.ShouldBind(&userInfo); err == nil {
		userName := strings.TrimSpace(userInfo.UserName)
		user := DB.Where("user_name = ?", userName).First(&users)
		if user.Error != nil {
			// 处理错误
			fmt.Println(user.Error)
			return
		}
		//for _, user := range users {
		//	fmt.Printf("ID: %d, Name: %s,", user.Id, user.UserName)
		//}
		if utils.MD5(users.Password) == utils.MD5(userInfo.Password) {
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
