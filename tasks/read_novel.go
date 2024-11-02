package tasks

import (
	"fmt"

	"github.com/DnullP/GraphLang-GO/db"
	"github.com/DnullP/GraphLang-GO/model/task"
)

func ReadNoval(text string) {
	db.Init("neo4j://localhost", "neo4j", "1331233456")

	//提取文本实体
	entities := task.ExtractEntities(text)
	nodes := make([]db.Node, 0)

	nameList := make([]string, 0)
	name2tag := make(map[string]string)
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

		name2tag[name] = tag
		nameList = append(nameList, name)
	}
	db.CreateNode(&nodes)

	// 提取实体关系
	relations := task.ExtractRelations(string(text), nameList)

	relationDVOList := make([]db.Relationship, 0)
	for _, relation := range relations {
		if name2tag[relation.Obj] == "" || name2tag[relation.Sub] == "" {
			continue
		}
		relationDVO := db.Relationship{
			FromName:    relation.Sub,
			FromTag:     name2tag[relation.Sub],
			ToName:      relation.Obj,
			ToTag:       name2tag[relation.Obj],
			Description: relation.Description,
			Relation:    relation.Rel,
		}
		relationDVOList = append(relationDVOList, relationDVO)
	}

	// 存入数据库
	db.CreateRelation(&relationDVOList)
}
