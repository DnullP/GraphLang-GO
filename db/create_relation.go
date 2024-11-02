package db

import (
	"context"
	"fmt"
	"log"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type Relationship struct {
	FromName    string
	FromTag     string
	ToName      string
	ToTag       string
	Description string
	Relation    string
}

func CreateRelation(relations *[]Relationship) {

	session := neo4jDB.NewSession(context.Background(), neo4j.SessionConfig{
		AccessMode: neo4j.AccessModeWrite, // 写模式
	})
	defer func() {
		if err := session.Close(context.Background()); err != nil {
			log.Fatalf("无法关闭会话: %v", err)
		}
	}()

	for _, rel := range *relations {
		err := addRelationship(session, rel)
		if err != nil {
			log.Fatalf("添加关系失败: %v", err)
		}
		fmt.Printf("成功添加关系: %s -> %s\n", rel.FromName, rel.ToName)
	}

	fmt.Println("所有关系已成功添加到 Neo4j 数据库。")

}

func addRelationship(session neo4j.SessionWithContext, rel Relationship) error {
	_, err := session.ExecuteWrite(context.Background(), func(tx neo4j.ManagedTransaction) (interface{}, error) {
		query := fmt.Sprintf(`
            MATCH (a:`+"`%s`"+` {name: $fromName, tag: $fromTag})
            MATCH (b:`+"`%s`"+` {name: $toName, tag: $toTag})
            CREATE (a)-[:`+"`%s`"+` {description: $description}]->(b)
            RETURN a, b
        `, rel.FromTag, rel.ToTag, rel.Relation)

		params := map[string]interface{}{
			"fromName":    rel.FromName,
			"fromTag":     rel.FromTag,
			"toName":      rel.ToName,
			"toTag":       rel.ToTag,
			"description": rel.Description,
		}

		result, err := tx.Run(context.Background(), query, params)
		if err != nil {
			return nil, err
		}

		if result.Next(context.Background()) {
			fromNode, _ := result.Record().Get("a")
			toNode, _ := result.Record().Get("b")
			fmt.Printf("创建关系: %v KNOWS %v\n", fromNode, toNode)
		}

		return nil, result.Err()
	})

	return err
}
