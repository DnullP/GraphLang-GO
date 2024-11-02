package task_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/DnullP/GraphLang-GO/model/task"
)

func TestExtracteRelation(t *testing.T) {
	text, ok := os.ReadFile("../../text.txt")
	if ok != nil {
		fmt.Println(ok.Error())
	}
	entities := []string{}
	for _, entity := range task.ExtractEntities(string(text)) {
		entities = append(entities, entity.(map[string]interface{})["name"].(string))
	}
	fmt.Println(entities)

	relationList := task.ExtractRelations(string(text), entities)

	for _, relation := range relationList {
		fmt.Println(relation)
	}
}
