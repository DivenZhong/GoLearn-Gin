package model

import (
	"GinDemo_v1/global"
)

type Recipe struct {
	global.BaseModel
	Name        string   `json:"name"         bson:"name"`
	Tags        []string `json:"tags"         bson:"tags"`
	Ingredients []string `json:"ingredients"  bson:"ingredients"`
	Instruction string   `json:"instruction" bson:"instruction"`
}
