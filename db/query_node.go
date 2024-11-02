package db

import (
	"context"
	"fmt"
	"log"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func QueryNearbyNode(name string) {

}

func QueryNodeWithTag(tag string) []string {
	session := neo4jDB.NewSession(context.Background(), neo4j.SessionConfig{
		AccessMode: neo4j.AccessModeRead,
	})
	defer func() {
		if err := session.Close(context.Background()); err != nil {
			log.Fatalf("无法关闭会话: %v", err)
		}
	}()
	resultsNode, err := getNodesByLabel(session, tag)
	if err != nil {
		panic(err)
	}
	names := make([]string, 0)
	for _, node := range resultsNode {
		names = append(names, node.Props["name"].(string))
	}
	return names
}

type NodeDTO struct {
	Name        string
	Description []string
	Tag         string
}

func QueryNodeWithName(name string) ([]NodeDTO, error) {
	session := neo4jDB.NewSession(context.Background(), neo4j.SessionConfig{
		AccessMode: neo4j.AccessModeRead,
	})
	defer func() {
		if err := session.Close(context.Background()); err != nil {
			log.Fatalf("无法关闭会话: %v", err)
		}
	}()
	result, err := session.ExecuteRead(context.Background(), func(tx neo4j.ManagedTransaction) (any, error) {
		// 构建 Cypher 查询
		query := `
            MATCH (n {name: $name})
            RETURN n.name AS name, n.tag AS tag, n.descriptions AS description
        `

		// 执行查询
		records, err := tx.Run(context.Background(), query, map[string]any{
			"name": name,
		})
		if err != nil {
			return nil, err
		}

		// 收集节点信息
		var nodes []NodeDTO
		for records.Next(context.Background()) {
			record := records.Record()
			name, _ := record.Get("name")
			tag, _ := record.Get("tag")

			description_, _ := record.Get("description")
			description := []string{}
			for _, d := range description_.([]interface{}) {
				description = append(description, d.(string))
			}

			nodes = append(nodes, NodeDTO{Name: name.(string), Tag: tag.(string), Description: description})
		}
		// 检查迭代器错误
		if err = records.Err(); err != nil {
			return nil, err
		}

		return nodes, nil
	})
	if err != nil {
		return nil, err
	}
	return result.([]NodeDTO), nil
}

func getNodesByLabel(session neo4j.SessionWithContext, label string) ([]neo4j.Node, error) {
	result, err := session.ExecuteRead(context.Background(), func(tx neo4j.ManagedTransaction) (any, error) {
		// 构建 Cypher 查询
		query := fmt.Sprintf("MATCH (n:%s) RETURN n", label)

		// 执行查询
		records, err := tx.Run(context.Background(), query, nil)
		if err != nil {
			return nil, err
		}

		// 收集节点信息
		var nodes []neo4j.Node
		for records.Next(context.Background()) {
			record := records.Record()
			node, ok := record.Get("n")
			if ok {
				nodes = append(nodes, node.(neo4j.Node))
			}
		}
		// 检查迭代器错误
		if err = records.Err(); err != nil {
			return nil, err
		}

		return nodes, nil
	})

	if err != nil {
		return nil, err
	}

	return result.([]neo4j.Node), nil
}
