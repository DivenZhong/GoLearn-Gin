package initialize

import (
	"GinDemo_v1/global"
)

func InitConfig() {
	// 初始化全局配置文件
	InitViperConfig()
	// 初始化zap日志
	InitLogger()
	// 初始化gorm
	InitGorm()
	// 初始化redis
	InitRedis()
	// 初始化MongoDb
	initMongoDb()
	// 初始化es
	InitES()
	// 捕获启动时错误
	defer global.CatchError()
	// 程序退出前释放资源
	//defer CloseResource()
}

// 程序退出前释放资源
func CloseResource() {
	// 关闭数据库连接
	if global.GvaMysqlClient != nil {
		db, _ := global.GvaMysqlClient.DB()
		_ = db.Close()
	}
}
