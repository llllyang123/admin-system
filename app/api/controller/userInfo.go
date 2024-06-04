package apiController

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"user-admin/app/admin"
	model2 "user-admin/app/admin/model"
	"user-admin/utils"
)

func init() {

}

func EditUserInfo(c *gin.Context) (bool, string) {
	type UserInfo struct {
		UserName string `json:"username" form:"username"`
		NickName string `json:"nickname" form:"nickname"`
		Phone    string `json:"phone" form:"phone"`
		Sex      string `json:"sex" form:"sex"`
	}
	newUserInfo := &UserInfo{}
	if err := c.ShouldBind(&newUserInfo); err == nil {
		db := admin.DB
		var model model2.UserInfo
		result := db.Where("user_name = ?", newUserInfo.UserName).First(&model)
		if result.Error != nil {
			// 如果没有找到用户，返回错误
			return false, "未找到该用户"
		}
		if newUserInfo.NickName != "" {
			model.NickName = newUserInfo.NickName
		}
		if newUserInfo.Phone != "" {
			model.Phone = newUserInfo.Phone
		}
		if newUserInfo.Sex != "" {
			model.Sex = newUserInfo.Sex
		}
		if err := db.Save(&model).Error; err != nil {
			return false, "保存失败"
		}
	} else {
		return false, err.Error()
	}
	return true, ""
}

func GetUserSalaryInfo(c *gin.Context) (model2.UserInfo, bool) {
	getUserId := c.Query("job_id")
	var userMode model2.UserInfo
	db := utils.DB
	if getUserId != "" {
		userInfo := db.Where("job_id = ? AND is_delete = ?", getUserId, 0).First(&userMode)
		fmt.Println("userMode", userMode)
		if userInfo.Error == nil {
			return userMode, false
		} else {
			return userMode, true
		}
	} else {
		return userMode, false
	}

}
