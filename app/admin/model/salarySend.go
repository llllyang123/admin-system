package model

import (
	"time"
)

type SalarySend struct {
	Id    uint `gorm:"primary_key;AUTO_INCREMENT"`
	JobId uint `gorm:"comment:员工工号id"`
	//Round       uint      `gorm:"comment:"`
	Taxation    float64   `gorm:"type:decimal(13,2); comment:纳税金额;default:0.00"`
	OvertimePay float64   `gorm:"type:decimal(13,2);comment:加班费;default:0.00"`
	Subsidy     float64   `gorm:"type:decimal(13,2);comment:补贴;default:0.00"`
	BasicSalary float64   `gorm:"type:decimal(13,2);comment:基本工资;default:0.00"`
	Salary      float64   `gorm:"type:decimal(13,2);comment:应发工资总额;default:0.00"`
	Deduction   float64   `gorm:"type:decimal(13,2);comment:扣除金额数;default:0.00"`
	Received    float64   `gorm:"type:decimal(13,2);comment:实际所得工资;default:0.00"`
	Month       uint      `gorm:"comment:工资所发月份"`
	CreatTime   time.Time `gorm:"autoCreateTime;comment:创建时间"`
	UpdateTime  time.Time `gorm:"autoUpdateTime;comment:更新时间"`
}

func (SalarySend) TableName() string {
	return "salary_send"
}
