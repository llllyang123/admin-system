package controller

import (
	"user-admin/app/admin/model"
)

type UserInfo interface {
	GET()
	ADD()
	Update()
	Delete()
}

type UserInfoModel struct {
	model.UserInfo
}

func (g UserInfoModel) GET() []string {

	return []string{}
}
