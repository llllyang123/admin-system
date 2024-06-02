package model

import (
	"time"
)

func init() {

}

type UserInfo struct {
	Id        uint `gorm:"primaryKey"`
	UserName  string
	NickName  string
	Email     string
	Password  string
	Salt      string
	Phone     string
	PhoneCode string
	Sex       string
	CreatTime time.Time `gorm:"autoCreateTime"`
	Uptime    time.Time `gorm:"autoUpdateTime"`
}

type UserInfoGet struct {
	Id        uint
	UserName  string
	NickName  string
	Email     string
	Sex       string
	CreatTime time.Time
	Uptime    time.Time
}

func (UserInfo) TableName() string {
	return "users"
}

func (UserInfoGet) TableName() string {
	return "users"
}
