package main

import (
	"HiChat/initialize"
	"HiChat/router"
)

func main() {
	initialize.InitConfig()
	//初始化数据库
	initialize.InitDB()
	initialize.InitRedis()
	//初始化日志
	initialize.InitLogger()

	route := router.Route()
	err := route.Run(":8000")
	if err != nil {
		return
	}
}
