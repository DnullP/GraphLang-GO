package model_test

import (
	"os"
	"testing"

	"github.com/DnullP/GraphLang-GO/model"
)

func TestModel(t *testing.T) {
	model := model.GlobelModel

	//prompt_template := "你的任务是文本中提取重要人物、地点、组织、物品、概念等对象，返回格式为：{\"obj\": [{\"name\": \"xxx\", \"type\": \"xxx\"}, ...]}：只返回json表即可，不要输出其他内容type字段应该尽量简短。"

	prompt_template_2 := "你的任务是提取文本中\"鵺\"和\"圣白莲\"的主客观描述或言行事迹，生成如下json格式的数据：{\"鵺\": [\"xxx\", \"xxx\", ...], \"圣白莲\": [\"xxx\", \"xxx\", ...]}, 包含多个对象的句子应该分别提取到对应对象的列表中"

	text, err := os.ReadFile("../text.txt")

	if err != nil {
		t.Error(err)
		return
	}

	result := model.Input(string(text) + prompt_template_2)
	t.Logf("%v", result)
	t.Log("test finish!")
}
