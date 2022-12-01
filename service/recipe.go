/**
 * @Description TODO
 **/
package service

import (
	"GinDemo_v1/global"
	"GinDemo_v1/model"
	"GinDemo_v1/model/request"
	"GinDemo_v1/model/response"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

/**
 * @description: 添加新菜谱
 * @param recipe
 */
func AddNewRecipe(ctx *gin.Context, recipeParam request.RecipeParam) (*model.Recipe, error) {
	recipe := model.Recipe{
		Name:        recipeParam.Name,
		Tags:        recipeParam.Tags,
		Ingredients: recipeParam.Ingredients,
		Instruction: recipeParam.Instruction,
	}
	collection := global.GvaMongoDbClient.Database("gin_demo").Collection("recipe")
	_, err := collection.InsertOne(ctx, recipe)
	if err != nil {
		response.Error(ctx, "添加新菜谱失败!")
		return &recipe, err
	}
	defer collection.Clone()

	return &recipe, nil
}

/**
 * @description: 获取菜谱列表
 * @param recipe
 */
func GetRecipeList(ctx *gin.Context) (*[]model.Recipe, error) {
	recipes := make([]model.Recipe, 0)
	collection := global.GvaMongoDbClient.Database("gin_demo").Collection("recipe")
	list, err := collection.Find(ctx, bson.M{})
	if err != nil {
		response.Error(ctx, "获取菜谱列表失败!")
		return &recipes, err
	}
	defer list.Close(ctx)
	defer collection.Clone()
	for list.Next(ctx) {
		var recipe model.Recipe
		list.Decode(&recipe)
		recipes = append(recipes, recipe)
	}

	return &recipes, nil
}
