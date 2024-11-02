package db

import (
	"context"
	"fmt"
	"log"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type Node struct {
	Name        string
	Description []string
	Tag         string
}

func CreateNode(nodes *[]Node) {

	session := neo4jDB.NewSession(context.Background(), neo4j.SessionConfig{
		AccessMode: neo4j.AccessModeWrite, // 写模式
	})
	defer func() {
		if err := session.Close(context.Background()); err != nil {
			log.Fatalf("无法关闭会话: %v", err)
		}
	}()

	for _, node := range *nodes {
		err := addNode(session, node)
		if err != nil {
			log.Fatalf("添加节点失败: %v", err)
		}
		fmt.Printf("成功添加节点: %s\n", node.Name)
	}

}

func addNode(session neo4j.SessionWithContext, node Node) error {

	_, err := session.ExecuteWrite(context.Background(), func(tx neo4j.ManagedTransaction) (interface{}, error) {

		query := fmt.Sprintf(`
        MERGE (n:%s {name: $name, tag: $tag})
            ON CREATE SET n.descriptions = $descriptions
            ON MATCH SET n.descriptions = coalesce(n.descriptions, []) + $newDescriptions
            RETURN n`, node.Tag)

		params := map[string]interface{}{
			"name":            node.Name,
			"descriptions":    node.Description,
			"newDescriptions": node.Description,
			"tag":             node.Tag,
		}

		result, err := tx.Run(context.Background(), query, params)
		if err != nil {
			return nil, err
		}

		if result.Next(context.Background()) {
			createdNode, _ := result.Record().Get("n")
			fmt.Printf("创建的节点: %v\n", createdNode)
		}

		return nil, result.Err()
	})

	return err
}
