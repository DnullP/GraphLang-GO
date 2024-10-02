package db_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/DnullP/GraphLang-GO/db"
	"github.com/DnullP/GraphLang-GO/model/task"
)

func TestCreateNode(t *testing.T) {
	db.Init("neo4j://localhost", "neo4j", "1331233456")

	text, err := os.ReadFile("../text.txt")
	if err != nil {
		panic(err)
	}
	entities := task.ExtractEntities(string(text))
	nodes := make([]db.Node, 0)

	for _, entity := range entities {

		name := entity.(map[string]interface{})["name"].(string)
		tag := entity.(map[string]interface{})["type"].(string)
		fmt.Printf("name: %s, tag: %s\n", name, tag)

		infos := task.ExtractInfomation(string(text), name)
		node := db.Node{
			Name:        name,
			Tag:         tag,
			Description: infos,
		}
		nodes = append(nodes, node)
	}
	db.CreateNode(&nodes)
}
