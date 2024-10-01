package task

import (
	"fmt"
	"strings"

	"github.com/DnullP/GraphLang-GO/model"
)

var prompt_relation_1 string = "你的任务是从上述文本中提取出关于下面对象实体之间的关系"
var prompt_relation_2 string = "并以{\"relation\": [{\"sub\":\"obj1\", \"obj\":\"obj2\", \"description\": \"sth\"}, ...]}的json格式返回, 你只需要输出json即可, 不需要其他任何内容"

func ExtractRelations(text string, entities []string) map[string]interface{} {
	jsonRaw := model.GlobelModel.Input(text + prompt_relation_1 + strings.Join(entities, "\n") + "\n" + prompt_relation_2)
	jsonRaw = removeFirstAndLastLine(jsonRaw)

	fmt.Println(jsonRaw)

	return nil
}
