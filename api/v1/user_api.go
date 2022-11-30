/**
 * @Description 用户相关表
 **/
package v1

import (
	"GinDemo_v1/global"
	"GinDemo_v1/middleware"
	"GinDemo_v1/model"
	"GinDemo_v1/model/request"
	"GinDemo_v1/model/response"
	"GinDemo_v1/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// Register
// @Summary  用户注册
// @Tags  用户
// @Accept json
// @Produce json
// @Param data body request.RegisterParam true "json数据"
// @Success 200 {object} model.User "{"code": 0,"msg": "请求成功", "data": {}}"
// @Router /v1/user/register [post]
func Register(ctx *gin.Context) {
	// 绑定参数
	var registerParam request.RegisterParam
	_ = ctx.ShouldBindJSON(&registerParam)
	// todo 参数校验
	// 调用注册
	register, err := service.Register(registerParam)
	if err != nil {
		response.Error(ctx, "注册失败!")
		return
	}
	response.OkWithData(ctx, register)
}

// Login
// @Summary  用户账号密码登录
// @Tags  用户
// @Accept json
// @Produce json
// @Param data body request.LoginParam true "json数据"
// @Success 200 {object} model.User "{"code": 0,"msg": "请求成功", "data": {}}"
// @Router /v1/user/login [post]
func Login(ctx *gin.Context) {
	// 绑定参数
	var loginParam request.LoginParam
	_ = ctx.ShouldBindJSON(&loginParam)
	if loginParam.Password == "" || loginParam.Phone == "" {
		response.Error(ctx, "手机号和密码不能为空！")
		return
	}
	// 调用登录服务
	userRecord := model.User{Phone: loginParam.Phone, Password: loginParam.Password}
	if err := service.LoginPwd(&userRecord); err != nil {
		global.GvaLogger.Error("登录失败:", zap.Any("user", userRecord))
		response.Error(ctx, "登录失败,账号或者密码错误!")
		return
	}
	// 生成token
	token, err := middleware.CreateToken(userRecord.ID)
	if err != nil {
		global.GvaLogger.Sugar().Errorf("登录失败,Token生成异常:%s", err)
		response.Error(ctx, "登录失败,账号或者密码错误!")
		return
	}
	userRecord.Token = token
	response.OkWithData(ctx, userRecord)
}

// GetUser
// @Summary  查询用户信息
// @Tags  用户
// @Accept json
// @Produce json
// @Success 200 {string} GetUser "{"code": 0,"msg": "请求成功", "data": {}}"
// @Router /v1/user/detail [get]
func GetUser(ctx *gin.Context) {
	// 从上下文中获取用户信息，(经过中间件逻辑时，已经设置到上下文)
	user, _ := ctx.Get("user")
	response.OkWithData(ctx, user)
	return
}
