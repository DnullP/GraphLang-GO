package tasks

import (
	"fmt"
	"log"

	"github.com/DnullP/GraphLang-GO/db"
	"github.com/DnullP/GraphLang-GO/model/task"
)

func MergeEntites(entities []string) {
	results := task.MergeEntities(entities)

	for baseName, nameSets := range results {
		fmt.Println(baseName, nameSets)

		//选取的名字有时候不在对象集合中
		inSet := false
		for _, name := range nameSets.([]interface{}) {
			if baseName == name.(string) {
				inSet = true
				break
			}
		}
		if !inSet {
			continue
		}

		for _, name := range nameSets.([]interface{}) {
			if name.(string) == baseName {
				continue
			}

			db.MergeNode(baseName, name.(string))
			log.Println(baseName, name)
		}
	}
}

func TryMergePerson() {
	db.Init("neo4j://localhost", "neo4j", "1331233456")
	persons := db.QueryNodeWithTag("人物")
	MergeEntites(persons)
}
