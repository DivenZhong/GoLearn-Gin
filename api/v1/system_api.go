/**
 * @Author Mr.ZhongQiWen
 * @Description 系统接口
 * @Date 2022/11/6 4:43 下午
 **/
package v1

import (
	"GinDemo_v1/global"
	"GinDemo_v1/model/response"
	"github.com/gin-gonic/gin"
)

// 配置信息
func GetConfig(ctx *gin.Context) {
	response.OkWithData(ctx, global.GvaConfig)
}
