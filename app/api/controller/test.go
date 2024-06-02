package apiController

import (
	"fmt"
	"user-admin/app/admin"
	"user-admin/app/admin/model"
)

func Test() ([]model.UserInfoGet, int64, error) {
	db := admin.DB
	var model []model.UserInfoGet

	userInfo := db.Where("user_name=?", "admin").First(&model)
	if userInfo.Error != nil {
		// 处理错误
		fmt.Println(userInfo.Error)
		return model, 0, nil
	}
	count := userInfo.RowsAffected
	return model, count, userInfo.Error

}
