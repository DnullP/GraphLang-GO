package db

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func main_neo() {
	ctx := context.Background()
	// URI examples: "neo4j://localhost", "neo4j+s://xxx.databases.neo4j.io"
	dbUri := "neo4j://localhost"
	dbUser := "neo4j"
	dbPassword := "1331233456"

	driver, err := neo4j.NewDriverWithContext(
		dbUri,
		neo4j.BasicAuth(dbUser, dbPassword, ""))
	if err != nil {
		panic(err)
	}
	defer driver.Close(ctx)

	session := driver.NewSession(context.Background(), neo4j.SessionConfig{
		AccessMode: neo4j.AccessModeWrite,
	})
	defer func() {
		if err = session.Close(context.Background()); err != nil {
			log.Fatalf("无法关闭会话: %v", err)
		}
	}()

	createResult, err := session.ExecuteWrite(context.Background(), func(tx neo4j.ManagedTransaction) (interface{}, error) {
		result, err := tx.Run(context.Background(),
			"CREATE (a:Person {name: $name}) RETURN a",
			map[string]interface{}{
				"name": "Alice",
			},
		)
		if err != nil {
			return nil, err
		}
		if result.Next(context.Background()) {
			node, ok := result.Record().Get("a")
			if ok {
				return node, nil
			}
		}
		return nil, result.Err()
	})

	if err != nil {
        log.Fatalf("创建节点失败: %v", err)
    }

    fmt.Printf("创建的节点: %v\n", createResult)

}

func TestNeo(t *testing.T) {
	main_neo()
}
