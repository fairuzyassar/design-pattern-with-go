package main

import "fmt"

type DBconnection struct{
	username, password string
}

var connector *DBconnection

func getConnection() *DBconnection {
	if connector == nil {
		connector = &DBconnection{"userdb", "passworddb"}
	}
	return connector
}

func main() {
	firstconnection := getConnection()
	secondconnection := getConnection()

	fmt.Println(firstconnection == secondconnection)
}