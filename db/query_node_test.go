package db_test

import (
	"fmt"
	"testing"

	"github.com/DnullP/GraphLang-GO/db"
)

func TestQueryNodeWithTag(t *testing.T) {
	db.Init("neo4j://localhost", "neo4j", "1331233456")
	fmt.Println(db.QueryNodeWithTag("人物"))
}

func TestQueryNodeWithName(t *testing.T) {	
	db.Init("neo4j://localhost", "neo4j", "1331233456")
	fmt.Println(db.QueryNodeWithName("岩永"))
}
