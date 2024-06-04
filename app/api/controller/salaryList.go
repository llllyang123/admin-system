package apiController

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	model2 "user-admin/app/admin/model"
	"user-admin/utils"
)

func GetSalaryList(c *gin.Context) ([]model2.SalarySend, int64, int64, int64) {
	var model []model2.SalarySend
	db := utils.DB
	// 获取查询参数
	pageStr := c.Query("page")
	pageSizeStr := c.Query("Limit")
	// 设置默认的分页参数
	page, err := strconv.Atoi(pageStr)
	if err != nil || page <= 0 {
		page = 1
	}
	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil || pageSize <= 0 {
		pageSize = 10
	}
	// 计算offset

	offset := (page - 1) * pageSize
	fmt.Println("pageSize: ", pageSize, "\n offset: ", offset)
	dbInfo := db.Limit(pageSize).Offset(offset).Find(&model)
	// 假设你想返回总记录数以供前端进行分页显示
	if dbInfo.Error != nil {
		fmt.Println("dbInfo.Error: ", dbInfo.Error)
	}
	count := dbInfo.RowsAffected
	return model, count, int64(page), int64(pageSize)
}

func AddSalaryList(c *gin.Context) (bool, error, int64) {
	type param struct {
		//UserId uint `json:"user_id"`
		//Round       float64 `json:"round"`
		JobId       uint    `json:"job_id" form:"job_id"`
		Taxation    float64 `json:"taxation" form:"taxation"`
		OvertimePay float64 `json:"overtime_pay" form:"overtime_pay"`
		Subsidy     float64 `json:"subsidy" form:"subsidy"`
		BasicSalary float64 `json:"basic_salary" form:"basic_salary"`
		Salary      float64 `json:"salary" form:"salary"`
		Deduction   float64 `json:"deduction" form:"deduction"`
		Received    float64 `json:"received" form:"received"`
		Month       uint    `json:"month" form:"month"`
	}
	Params := &param{}
	if err := c.ShouldBind(&Params); err == nil {
		var userModel model2.UserInfoGet
		db := utils.DB
		isUser := db.Where("job_id = ?", Params.JobId).Find(&userModel)
		if isUser.Error != nil {
			fmt.Println("isUser.Error: ", isUser.Error)
			return false, isUser.Error, 2
		}
		if isUser == nil {
			fmt.Println("userModel: ", userModel, "\n 该用户查询不到")
			return false, isUser.Error, 1
		}
		var model model2.SalarySend
		getInfo := db.Where("job_id = ? AND month = ?", Params.JobId, Params.Month).Find(&model)
		if getInfo.Error != nil {
			fmt.Println("getInfo.Error: ", getInfo.Error)
			return false, getInfo.Error, 2
		}
		fmt.Printf("getInfo: %d\n", getInfo)
		fmt.Printf("userModel: %d", userModel)
		if getInfo.RowsAffected != 0 {
			fmt.Println("getInfo: 已经有该数据")
			return false, nil, 1
		}
		//model.UserId = Params.UserId
		model.JobId = Params.JobId
		model.Taxation = Params.Taxation
		model.OvertimePay = Params.OvertimePay
		model.Subsidy = Params.Subsidy
		model.Deduction = Params.Deduction
		model.BasicSalary = Params.BasicSalary
		model.Salary = Params.Salary
		model.Month = Params.Month
		request := db.Create(&model)
		if request.Error != nil {
			fmt.Println("request.Error: ", request.Error)
			return false, request.Error, 2
		}
		return true, nil, 0
	} else {
		fmt.Println(err)
		return false, err, 2
	}

}
