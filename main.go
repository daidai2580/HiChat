package main

import (
	"HiChat/initialize"
	"HiChat/router"
)

func main() {
	//初始化数据库
	initialize.InitDB()
	//初始化日志
	initialize.InitLogger()

	route := router.Route()
	err := route.Run(":8000")
	if err != nil {
		return
	}
}
