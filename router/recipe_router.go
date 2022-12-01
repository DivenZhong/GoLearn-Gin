package router

import (
	v1 "GinDemo_v1/api/v1"
	"github.com/gin-gonic/gin"
)

/**
 * @description: 菜谱相关的路由
 * @param engine
 */
func InitRecipeRouter(engine *gin.Engine) {
	// 不需要登录的路由
	recipeRouter := engine.Group("v1/recipe")
	{
		// 添加新菜谱
		recipeRouter.POST("add", v1.AddNewRecipe)
		// 获取菜谱列表
		recipeRouter.POST("list", v1.GetRecipeList)
	}
}
