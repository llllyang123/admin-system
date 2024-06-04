package model

import (
	"time"
)

func init() {

}

type UserInfo struct {
	Id          uint      `gorm:"primary_key;AUTO_INCREMENT"`
	JobId       uint      `gorm:"primary_key;default:0;AUTO_INCREMENT"`
	UserName    string    `gorm:"comment:用户名"`
	NickName    string    `gorm:"comment:昵称"`
	Avatar      string    `gorm:"comment:头像"`
	Email       string    `gorm:"comment:email"`
	Password    string    `gorm:"comment:password"`
	Salt        string    `gorm:"comment:salt"`
	Phone       string    `gorm:"comment:phone"`
	PhoneCode   string    `gorm:"comment:phone_code;"`
	Sex         string    `gorm:"comment:性别"`
	BasicSalary float64   `gorm:"comment:基本工资;default:0.00"`
	Salary      float64   `gorm:"comment:工资;default:0.00"`
	Rank        uint      `gorm:"comment:职级;default:0"`
	IsDelete    uint      `gorm:"comment:是否删除，0=否，1=是;default:0"`
	CreatTime   time.Time `gorm:"autoCreateTime;comment:创建时间"`
	Uptime      time.Time `gorm:"autoUpdateTime;comment:更新时间"`
}

func (UserInfo) TableName() string {
	return "users"
}

type UserInfoGet struct {
	Id        uint
	JobId     uint      `gorm:"primary_key;AUTO_INCREMENT"`
	UserName  string    `gorm:"comment:user_name"`
	NickName  string    `gorm:"comment:nick_name"`
	Avatar    string    `gorm:"comment:avatar"`
	Email     string    `gorm:"comment:email"`
	Sex       string    `gorm:"comment:sex"`
	IsDelete  uint      `gorm:"comment:是否删除，0=否，1=是"`
	CreatTime time.Time `gorm:"autoCreateTime;comment:创建时间"`
	Uptime    time.Time `gorm:"autoUpdateTime;comment:更新时间"`
}

func (UserInfoGet) TableName() string {
	return "users"
}
