package db

import "fmt"

type DbNoSql struct {
	DbBase
}

type NoSqlClient struct {
	DatabaseName string
}

func (db *DbNoSql) Connect() bool {
	client := new(NoSqlClient)

	client.DatabaseName = "MongoDB"

	db.Connection = client

	fmt.Println("Connecting from a No SQL batabase")

	return true
}

func (db *DbNoSql) DbClientName() string {
	client := db.Connection.(*NoSqlClient)

	return client.DatabaseName
}
