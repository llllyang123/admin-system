package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"user-admin/app/admin/model"
	"user-admin/utils"
)

func UserInfo() (model.UserInfo, bool) {
	var userMode model.UserInfo
	db := utils.DB
	userId := utils.Session.Get("userId")
	if userId != "" {
		userInfo := db.Where("id = ?", userId).First(&userMode)
		if userInfo.Error == nil {
			return userMode, true
		} else {
			return userMode, false
		}
	} else {
		return userMode, false
	}

}

func UpdateUserInfo(c *gin.Context) (model.UserInfo, string) {
	type UserInfo struct {
		UserName string `json:"username" form:"username"`
		NickName string `json:"nickname" form:"nickname" default:""`
		Sex      string `json:"sex" form:"sex" default:""`
		Avatar   string `json:"avatar" form:"avatar"`
		Phone    string `json:"phone" form:"phone" default:""`
		Email    string `json:"email" form:"email" default:""`
	}
	newUserInfo := &UserInfo{}

	if err := c.ShouldBind(&newUserInfo); err == nil {
		db := utils.DB
		userId := utils.Session.Get("userId")
		var model model.UserInfo
		result := db.Where("id = ?", userId).First(&model)
		//result := db.First(&model, userId)
		if result.Error != nil {
			// 如果没有找到用户，返回错误
			return model, "未找到用户"
		}
		fmt.Println("user: ", model)
		if newUserInfo.NickName != "" {
			model.NickName = newUserInfo.NickName
		}
		if newUserInfo.Phone != "" {
			model.Phone = newUserInfo.Phone
		}
		if newUserInfo.Sex != "" {
			model.Sex = newUserInfo.Sex
		}
		if newUserInfo.Email != "" {
			model.Email = newUserInfo.Email
		}
		if newUserInfo.Avatar != "" {
			model.Sex = newUserInfo.Avatar
		}
		if err := db.Save(&model).Error; err != nil {
			return model, "保存失败"
		}
	} else {
		return model.UserInfo{}, ""
	}
	return model.UserInfo{}, ""
}
