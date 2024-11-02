package task

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/DnullP/GraphLang-GO/model"
)

var prompt_relation_1 string = "你的任务是从上述文本中提取出关于下面对象实体之间的关系, 关系不超过4个字, 关系描述不超过15个字, 每个对象可作为主体或者客体"
var prompt_relation_2 string = `并以{"relation": [{"rel":"rel_content","sub":"obj1", "obj":"obj2", "description": "sth"}, ...]}的json格式返回, 你只需要输出json即可, 不需要其他任何内容`

type Relation struct {
	Sub         string
	Obj         string
	Rel         string
	Description string
}

/*
Returned data format is as below:
{
    "relation": [
        {
            "rel": "rel_content",
            "sub": "obj1",
            "obj": "obj2",
            "description": "sth"
        }, ...
    ]
}
*/
func ExtractRelations(text string, entities []string) []Relation {

	jsonRaw := model.GlobelModel.Input(text + prompt_relation_1 + strings.Join(entities, "\n") + "\n" + prompt_relation_2)
	jsonRaw = removeFirstAndLastLine(jsonRaw)
	var jsonData map[string]interface{}
	json.Unmarshal([]byte(jsonRaw), &jsonData)

	relations := jsonData["relation"].([]interface{})
	relationList := make([]Relation, 0)

	for _, relation := range relations {
		relationData := relation.(map[string]interface{})
		if relationData["rel"] == "" {
			fmt.Println("stop here")
		}
		relationList = append(relationList, Relation{
			Sub:         relationData["sub"].(string),
			Obj:         relationData["obj"].(string),
			Rel:         relationData["rel"].(string),
			Description: relationData["description"].(string),
		})
	}

	return relationList
}
