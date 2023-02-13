package main

import (
	"douyin/config"
	"douyin/dao"
	"douyin/router"
	"douyin/utils"
	"fmt"
)

func main() {
	// fmt.Println(utils.GetLocalIp())
	fmt.Println(utils.LocalIPv4())

	// 初始化数据库
	dao.InitDataBase()

	// 初始化路由
	router := router.InitRouter()
	router.Run(fmt.Sprintf(":%d", config.Info.Port))
}
