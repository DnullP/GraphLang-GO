package db_test

import (
	"os"
	"testing"

	"github.com/DnullP/GraphLang-GO/db"
	"github.com/DnullP/GraphLang-GO/model/task"
)

func TestCreateRelation(t *testing.T) {
	db.Init("neo4j://localhost", "neo4j", "1331233456")

	text, err := os.ReadFile("../text.txt")
	if err != nil {
		panic(err)
	}

	entities := task.ExtractEntities(string(text))

	nameList := make([]string, 0)
	name2tag := make(map[string]string)
	for _, entity := range entities {
		name := entity.(map[string]interface{})["name"].(string)
		tag := entity.(map[string]interface{})["type"].(string)
		name2tag[name] = tag
		nameList = append(nameList, name)
	}

	relations := task.ExtractRelations(string(text), nameList)

	relationDVOList := make([]db.Relationship, 0)
	for _, relation := range relations {
		relationDVO := db.Relationship{
			FromName:    relation.Sub,
			FromTag:     name2tag[relation.Sub],
			ToName:      relation.Obj,
			ToTag:       name2tag[relation.Obj],
			Description: relation.Description,
		}
		relationDVOList = append(relationDVOList, relationDVO)
	}
	db.CreateRelation(&relationDVOList)
}
