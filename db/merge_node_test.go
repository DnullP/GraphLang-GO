package db_test

import (
	"testing"

	"github.com/DnullP/GraphLang-GO/db"
)

func TestMergeNode(t *testing.T) {
	db.Init("neo4j://localhost", "neo4j", "1331233456")
	db.MergeNode("岩永", "岩永琴子")
}
