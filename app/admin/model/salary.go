package model

import (
	"time"
)

type Salary struct {
	Id          uint      `gorm:"primary_key;AUTO_INCREMENT"`
	UserId      uint      `gorm:"comment:用户表id"`
	OvertimePay float64   `gorm:"type:decimal(13,2);comment:加班费;default:0.00"`
	Subsidy     float64   `gorm:"type:decimal(13,2);comment:补贴;default:0.00"`
	BasicSalary float64   `gorm:"type:decimal(13,2);comment:基本工资;default:0.00"`
	salary      float64   `gorm:"type:decimal(13,2);comment:应发工资总额;default:0.00"`
	CreatTime   time.Time `gorm:"autoCreateTime;comment:创建时间"`
	UpdateTime  time.Time `gorm:"autoUpdateTime;comment:更新时间"`
}

func (Salary) TableName() string {
	return "salary"
}
