package main

import (
	"GinDemo_v1/core"
	"GinDemo_v1/initialize"
)

func main() {
	// 初始化配置
	initialize.InitConfig()
	// 启动服务
	core.RunServer()
}
