/**
 * @Description TODO
 **/
package v1

import (
	"GinDemo_v1/global"
	"GinDemo_v1/model/response"
	"github.com/gin-gonic/gin"
	"time"
)

// RdTest
// @Summary  验证redis
// @Tags  测试验证
// @Accept json
// @Produce json
// @Param type query string true "请求类型" maxlength(15)
// @Success 200 {string} RdTest "{"code": 0,"msg": "请求成功", "data": {}}"
// @Router /test/redis [get]
func RdTest(ctx *gin.Context) {
	method, _ := ctx.GetQuery("type")
	var result string
	var err error
	switch method {
	case "get":
		result, err = global.GvaRedis.Get(ctx, "test").Result()
	case "set":
		result, err = global.GvaRedis.Set(ctx, "test", "hello word!", time.Hour).Result()
	}
	if err != nil {
		response.Error(ctx, err.Error())
		return
	}
	response.OkWithData(ctx, result)
}
