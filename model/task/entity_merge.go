package task

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/DnullP/GraphLang-GO/model"
)

var prompt_entity_merge_1 string = "你的任务是将下述列表中明显是指代相同对象的名词合并返回:"
var prompt_entity_merge_2 string = "\n返回格式为：{\"base_name\": [\"name_1\", ...], ...},base_name是合并后的新名字, 数组中是被合并的名字,你只返回json表即可,对于无法完全确定的不用返回"

/*
Returned data format is below:

	{
		"base_name": ["name_1", ...],
		...
	}
*/
func MergeEntities(entities []string) map[string]interface{} {
	jsonRaw := model.GlobelModel.Input(prompt_entity_merge_1 + string(strings.Join(entities, " ")) + prompt_entity_merge_2)

	jsonRaw = removeFirstAndLastLine(jsonRaw)

	var jsonData map[string]interface{}
	ok := json.Unmarshal([]byte(jsonRaw), &jsonData)
	if ok != nil {
		fmt.Println(jsonRaw)
		panic("Model doesn't return a json")
	}
	return jsonData
}
