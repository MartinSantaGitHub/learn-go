package db

import (
	"fmt"
)

type DbSql struct {
	DbBase
}

type SqlClient struct {
	DbName string
}

func (db *DbSql) Connect() bool {
	client := new(SqlClient)

	client.DbName = "PostgreSQL"

	db.Connection = client

	fmt.Println("Connecting from a SQL batabase")

	return true
}

func (db *DbSql) DbClientName() string {
	client := db.Connection.(*SqlClient)

	return client.DbName
}
