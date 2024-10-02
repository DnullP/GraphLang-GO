package task_test

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/DnullP/GraphLang-GO/model/task"
)

func LogJsonData(data map[string]interface{}) {
	file, err := os.Create("log.json")
	if err != nil {
		panic("Could not create log file!")
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(data); err != nil {
		panic("Could not encode JSON data!")
	}
}

func TestExtract(t *testing.T) {
	text, ok := os.ReadFile("../text.txt")
	if ok != nil {
		panic("Reading file error!")
	}
	jsonData := task.ExtractEntities(string(text))
	for _, data := range jsonData {
		fmt.Println(data.(map[string]interface{})["name"].(string))
	}
}
