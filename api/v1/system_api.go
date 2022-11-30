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

// GetConfig
// @Summary  系统配置信息
// @Tags  系统管理
// @Accept json
// @Produce json
// @Success 200 {string} GetConfig "{"code": 0,"msg": "请求成功", "data": {}}"
// @Router /system/config [get]
func GetConfig(ctx *gin.Context) {
	response.OkWithData(ctx, global.GvaConfig)
}
