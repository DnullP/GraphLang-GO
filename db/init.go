package db

import "github.com/neo4j/neo4j-go-driver/v5/neo4j"

func Init(url string, user string, passwd string) {
	driver, err := neo4j.NewDriverWithContext(
		url,
		neo4j.BasicAuth(user, passwd, ""))
	if err != nil {
		panic(err)
	}
	neo4jDB = driver
}
