package request

// 菜谱参数
type RecipeParam struct {
	Name        string   `json:"name"`        //菜谱名称
	Tags        []string `json:"tags"`        //标签列表
	Ingredients []string `json:"ingredients"` //成分
	Instruction string   `json:"instruction"` //说明书
}
