package v1

import (
	"GinDemo_v1/model/request"
	"GinDemo_v1/model/response"
	"GinDemo_v1/service"
	"github.com/gin-gonic/gin"
)

// AddNewRecipe
// @Summary  添加新的菜谱
// @Tags  菜谱
// @Accept json
// @Produce json
// @Param data body request.RecipeParam true "json数据"
// @Success 200 {object} model.Recipe "{"code": 0,"msg": "请求成功", "data": {}}"
// @Router /v1/recipe/add [post]
func AddNewRecipe(ctx *gin.Context) {
	// 绑定参数
	var recipeParam request.RecipeParam
	_ = ctx.ShouldBindJSON(&recipeParam)
	// 调用添加新菜谱
	recipe, err := service.AddNewRecipe(ctx, recipeParam)
	if err != nil {
		response.Error(ctx, "添加新菜谱失败!")
		return
	}
	response.OkWithData(ctx, recipe)
}

// GetRecipeList
// @Summary  获取菜谱列表
// @Tags  菜谱
// @Accept json
// @Produce json
// @Success 200 {object} []model.Recipe "{"code": 0,"msg": "请求成功", "data": {}}"
// @Router /v1/recipe/list [post]
func GetRecipeList(ctx *gin.Context) {
	// 绑定参数
	var recipeParam request.RecipeParam
	_ = ctx.ShouldBindJSON(&recipeParam)
	// 调用获取菜谱列表
	recipes, err := service.GetRecipeList(ctx)
	if err != nil {
		response.Error(ctx, "获取菜谱列表失败!")
		return
	}
	response.OkWithData(ctx, recipes)
}
