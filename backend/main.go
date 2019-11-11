package main

import (
	"comment/config"
	"comment/router"
)

func main() {
	// 从配置文件读取配置
	config.Init()
	// 装载路由
	r := router.NewRouter()

	if err := r.Run(":8888"); err != nil {
		panic(err)
	}
}

