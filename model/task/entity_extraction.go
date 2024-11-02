package task

import (
	"encoding/json"
	"fmt"

	"github.com/DnullP/GraphLang-GO/model"
)

var prompt_entity string = "你的任务是文本中提取重要人物、地点、组织、物品和概念, 不要提取代词，返回格式为：{\"obj\": [{\"name\": \"xxx\", \"type\": \"xxx\"}, ...]},只返回json表即可,type字段应当为人物、地点、组织、物品、概念之一"

/*
Returned data format is below:

	{
		"obj": [
			{
				"name": "xxx",
				"type": "xxx"
			},
			...
		]
	}
*/
func ExtractEntities(text string) []interface{} {

	jsonRaw := model.GlobelModel.Input(text + prompt_entity)

	jsonRaw = removeFirstAndLastLine(jsonRaw)

	var jsonData map[string]interface{}
	ok := json.Unmarshal([]byte(jsonRaw), &jsonData)
	if ok != nil {
		fmt.Println(jsonRaw)
		panic("Model doesn't return a json")
	}

	entityList := make([]interface{}, 0)
	objList := jsonData["obj"].([]interface{})

	entityList = append(entityList, objList...)

	return entityList
}
