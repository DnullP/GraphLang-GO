package db

import (
	"context"
	"log"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func MergeNode(nameA string, nameB string) {
	session := neo4jDB.NewSession(context.Background(), neo4j.SessionConfig{
		AccessMode: neo4j.AccessModeRead,
	})
	defer func() {
		if err := session.Close(context.Background()); err != nil {
			log.Fatalf("无法关闭会话: %v", err)
		}
	}()

	_, err := session.ExecuteWrite(context.Background(), func(tx neo4j.ManagedTransaction) (any, error) {
		// 构建 Cypher 查询
		query1 := `
MATCH (a1{name:$nameA}), (a2{name:$nameB})
WITH head(collect([a1,a2])) as nodes
CALL apoc.refactor.mergeNodes(nodes,{properties:"discard", mergeRels:true})
YIELD node
RETURN count(*)
			  `
		// 执行查询
		_, err := tx.Run(context.Background(), query1, map[string]any{
			"nameA": nameA, "nameB": nameB,
		})
		if err != nil {
			return nil, err
		}
		return nil, nil
	})
	if err != nil {
		panic(err)
	}
}
