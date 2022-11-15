/**
 * @Author Mr.ZhongQiWen
 * @Description 路由注册入口
 * @Date 2022/11/5 3:17 下午
 **/
package core

import (
	"GinDemo_v1/router"
	"github.com/gin-gonic/gin"
)

// 注册路由入口
func RegisterRouters(engine *gin.Engine) {
	// 注册系统路由
	router.InitSystemRouter(engine)
	// 注册用户路由
	router.InitUserRouter(engine)
	// 测试路由
	router.InitTestRouter(engine)
	// 注册es路由
	router.InitESRouter(engine)
}