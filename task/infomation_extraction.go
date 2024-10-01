package task

import (
	"fmt"

	"github.com/DnullP/GraphLang-GO/model"
)

var prompt_infomation_1 = "你的任务是从上述文本中提取出对象\""
var prompt_infomation_2 = "\"相关的事实、描述、行为事迹等相关信息, 并以{\"name\":\"...\", \"info\":[\"...\", ...]}的json格式返回, 你只需要返回json, 不需要额外的文字内容"

func ExtractInfomation(text string, obj string) {
	jsonRaw := model.GlobelModel.Input(text + prompt_infomation_1 + obj + prompt_infomation_2)
	jsonRaw = removeFirstAndLastLine(jsonRaw)

	fmt.Println(jsonRaw)
}
