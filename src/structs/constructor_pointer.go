package main

import "fmt"

type DbConnection struct {
	isOpen   bool
	str_conn string
}

func NewDbConnection(str_conn string) *DbConnection {
	return &DbConnection{
		isOpen:   true,
		str_conn: str_conn,
	}
}

func (conn *DbConnection) Close() {
	conn.isOpen = false
	fmt.Println("Closing the DB conncetion ", conn.str_conn)
}

func main() {
	dbConn := NewDbConnection("**nice connection!***")
	performQuery1(dbConn)
	performQuery2(dbConn)
	performQuery3(dbConn)

	dbConn.Close()

	performQuery1(dbConn)
	performQuery2(dbConn)
	performQuery3(dbConn)
}

func performQuery1(conn *DbConnection) {
	fmt.Println("Query 1 executed", conn.str_conn, conn.isOpen)
}

func performQuery2(conn *DbConnection) {
	fmt.Println("Query 2 executed", conn.str_conn, conn.isOpen)
}

func performQuery3(conn *DbConnection) {
	fmt.Println("Query 3 executed", conn.str_conn, conn.isOpen)
}
