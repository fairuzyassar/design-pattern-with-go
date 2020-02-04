package main

import "fmt"

type Connection struct {
	id int
}

type DatabasePool struct {
	pool []*Connection
}

func (dp *DatabasePool) initalizationConnection(){
	for i:=0; i < 2; i++ {
		dp.pool = append(dp.pool, &Connection{i})
	}
}
func (dp *DatabasePool) acquireConnection() *Connection{
	var connection *Connection
	if(len(dp.pool) <= 0){
		fmt.Println("All connection locked")
	} else {
		connection = dp.pool[len(dp.pool)-1]
		dp.pool = dp.pool[:len(dp.pool)-1]
	}
	return connection
}

func (dp *DatabasePool) releaseConnection(connection *Connection) {
	dp.pool = append(dp.pool, connection)
}

func main(){
	databasePool := &DatabasePool{}
	databasePool.initalizationConnection()
	connection1 := databasePool.acquireConnection()
	fmt.Println(connection1.id)

	connection2 := databasePool.acquireConnection()
	fmt.Println(connection2.id)
	databasePool.releaseConnection(connection2)

	connection3 := databasePool.acquireConnection()
	fmt.Println(connection3.id)
}