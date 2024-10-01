package task

import (
	"encoding/json"
	"fmt"

	"github.com/DnullP/GraphLang-GO/model"
)

var prompt_entity string = "你的任务是文本中提取重要人物、地点、组织、物品、概念等对象，返回格式为：{\"obj\": [{\"name\": \"xxx\", \"type\": \"xxx\"}, ...]}：只返回json表即可，不要输出其他内容type字段应该尽量简短"

// das
func ExtractEntities(text string) []interface{} {

	jsonRaw := model.GlobelModel.Input(text + prompt_entity)

	jsonRaw = removeFirstAndLastLine(jsonRaw)

	var jsonData map[string]interface{}
	ok := json.Unmarshal([]byte(jsonRaw), &jsonData)
	if ok != nil {
		fmt.Println("Model doesn't return a json")
		return nil
	}

	entityList := make([]interface{}, 0)
	objList := jsonData["obj"].([]interface{})

	entityList = append(entityList, objList...)

	return entityList
}
