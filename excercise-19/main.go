package main

import (
	"db"
	"fmt"
)

var dbConn db.DbAdapter
var dbType string = "NoSQL"

func main() {
	dbConn = db.GetDataBase(dbType)
	isConnect := dbConn.Connect()

	fmt.Printf("Connected: %t\n", isConnect)
	fmt.Printf("Database Name: %s\n", dbConn.DbClientName())
	fmt.Println("End")
}
