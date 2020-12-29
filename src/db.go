package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

// Setup attempts to initialize the database and returns a handle for futher activity.
func Setup() {
	fmt.Println("Configuring DB")

	database, _ := sql.Open("sqlite3", "./todo.db")
	createDbStatement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS todos (id INTEGER PRIMARY KEY, title TEXT, done BOOLEAN)")
	createDbStatement.Exec()
}
