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
