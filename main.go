package main

import (
	"fmt"
	"user-admin/app/admin"
	"user-admin/app/api"
	"user-admin/app/index"
	"user-admin/routers"
	"user-admin/utils"
)

func main() {
	// 初始化redis init
	utils.RedisServiceInit()
	// 加载多个APP的路由配置
	routers.Include(admin.Routers, api.Routers, index.Routers)
	// 初始化路由
	r := routers.Init()

	//数据库连接
	// admin.Main()

	fmt.Println("Click open web: http://127.0.0.1:8001")
	// Listen and serve on 0.0.0.0:8080
	r.Run(":8001")
}
