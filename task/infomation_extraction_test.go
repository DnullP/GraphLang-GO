package task_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/DnullP/GraphLang-GO/task"
)

func TestExtractInfomation(t *testing.T) {
	text, ok := os.ReadFile("../text.txt")
	if ok != nil {
		fmt.Println(ok.Error())
	}
	entities := task.ExtractEntities(string(text))
	fmt.Println(entities...)
	for _, entity := range entities {
		name := entity.(map[string]interface{})["name"].(string)
		task.ExtractInfomation(string(text), name)
	}
}
