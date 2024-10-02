package task

import (
	"encoding/json"

	"github.com/DnullP/GraphLang-GO/model"
)

var prompt_infomation_1 = "你的任务是从上述文本中提取出对象\""
var prompt_infomation_2 = "\"相关的事实、性质、行为事迹等相关信息, 每条描述不超过15字, 其中名词而不是代词, 并以{\"name\":\"...\", \"info\":[\"...\", ...]}的json格式返回, 你只需要返回json, 不需要额外的文字内容"

/*
Returned data format is as below:

	{
		"name": "...",
		"info": [
			"....",
			...
		]
	}
*/
func ExtractInfomation(text string, obj string) []string {
	jsonRaw := model.GlobelModel.Input(text + prompt_infomation_1 + obj + prompt_infomation_2)
	jsonRaw = removeFirstAndLastLine(jsonRaw)

	var jsonData map[string]interface{}
	json.Unmarshal([]byte(jsonRaw), &jsonData)

	result := make([]string, 0)
	infos := jsonData["info"].([]interface{})
	for _, info := range infos {
		result = append(result, info.(string))
	}
	return result
}
