package task

import (
	"encoding/json"
	"strings"

	"github.com/DnullP/GraphLang-GO/model"
)

var prompt_similar_1 string = "你的任务是从下面的对象列表中, 将你认为是代指相同对象的名词选出来, 并返回结果:"
var prompt_similar_2 string = `你需要以{"sets": [["name_1", "name_2", ...], ["name_1_1", ...]]}的json格式返回, 然后你会得到每个集合中名词的具体描述, 用以进一步判断是否为同一对象, 所以你可以自由猜测;没有冗余指代的名词不用返回;你只需要返回json数据, 不需要其他内容`

func SelectSimilar(names []string) []interface{} {
	jsonRaw := model.GlobelModel.Input(prompt_similar_1 + strings.Join(names, "\n") + "\n" + prompt_similar_2)
	jsonRaw = removeFirstAndLastLine(jsonRaw)
	var jsonData map[string]interface{}
	json.Unmarshal([]byte(jsonRaw), &jsonData)

	sets := jsonData["sets"].([]interface{})
	return sets
}
