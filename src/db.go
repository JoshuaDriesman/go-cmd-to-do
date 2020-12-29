package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

// Todo represents a single todo item
type Todo struct {
	id    int
	title string
	done  bool
}

// Setup attempts to initialize the database and returns a handle for futher activity.
func Setup() *sql.DB {
	fmt.Println("Configuring DB")

	database, _ := sql.Open("sqlite3", "./todo.db")
	createStatement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS todos (id INTEGER PRIMARY KEY, title TEXT, done BOOLEAN)")
	createStatement.Exec()

	return database
}

// InsertTodo inserts a todo into the DB
func InsertTodo(title string, database *sql.DB) {
	insertStatement, _ := database.Prepare("INSERT INTO todos (title, done) VALUES (?, false)")
	insertStatement.Exec(title)
}

// GetTodos lists all of the todos in the DB
func GetTodos(database *sql.DB) *[]Todo {
	results := make([]Todo, 0)

	rows, _ := database.Query("SELECT * FROM todos")

	for rows.Next() {
		var result Todo
		rows.Scan(&result.id, &result.title, &result.done)
		results = append(results, result)
	}

	return &results
}

func ToggleTodo(id int, database *sql.DB) bool {
	updateStatement, _ := database.Prepare("UPDATE todos SET done = NOT done WHERE id=?")
	result, _ := updateStatement.Exec(id)

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected < 1 {
		return false
	}

	return true
}
