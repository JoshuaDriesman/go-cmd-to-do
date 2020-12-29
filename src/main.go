package main

import (
	"bufio"
	"database/sql"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var todoList = make(map[string]bool)

func checkAndHandleError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return
}

func listTodo(database *sql.DB) {
	tl := GetTodos(database)
	for _, todo := range *tl {
		fmt.Printf("%s: ", todo.title)
		if todo.done {
			fmt.Print("Completed\n")
		} else {
			fmt.Print("Incomplete\n")
		}
	}
}

func enumerateTodo() []string {
	enumeratedTodoItems := make([]string, len(todoList))
	index := 0
	for todoItem := range todoList {
		enumeratedTodoItems[index] = todoItem
		index++
	}
	return enumeratedTodoItems
}

func selectTodo() (string, error) {
	enumeratedTodoItems := enumerateTodo()
	for i, todoItem := range enumeratedTodoItems {
		fmt.Printf("%d: %s\n", i, todoItem)
	}
	fmt.Print("Enter number: ")

	var indexOfTodoToToggle string
	_, err := fmt.Scanln(&indexOfTodoToToggle)
	checkAndHandleError(err)

	intIndexOfTodoToToggle, err := strconv.Atoi(indexOfTodoToToggle)

	if err != nil || intIndexOfTodoToToggle > len(enumeratedTodoItems) {
		fmt.Println("You did not enter the number for a valid to do item!")
		return "", errors.New("Bad")
	}

	return enumeratedTodoItems[intIndexOfTodoToToggle], nil
}

func toggleTodo(database *sql.DB) {
	fmt.Println("Enter which to do item you'd like to toggle?")
	todos := GetTodos(database)
	for _, todo := range *todos {
		fmt.Println("%d: %s", todo.id, todo.title)
	}

	var input string
	fmt.Scanln(&input)

	num, _ := strconv.Atoi(input)

	ToggleTodo(num, database)

}

func deleteTodo() {
	fmt.Println("Enter which to do item you'd like to delete?")
	itemToDelete, err := selectTodo()
	if err != nil {
		return
	}
	delete(todoList, itemToDelete)
}

func promptAndRead(database *sql.DB) {
	for {
		fmt.Print("Enter a command (type 'h' for a list of commands): ")

		in := bufio.NewReader(os.Stdin)

		var command string

		_, err := fmt.Scanln(&command)

		checkAndHandleError(err)

		switch command {
		case "a":
			fmt.Print("Please enter the task you'd like to add to your to do list: ")

			newTodo, err := in.ReadString('\n')
			checkAndHandleError(err)

			newTodo = strings.Trim(newTodo, "\n")
			InsertTodo(newTodo, database)
		case "l":
			listTodo(database)
		case "t":
			toggleTodo(database)
		case "d":
			deleteTodo()
		case "h":
			fmt.Println("Here are all the commands this program supports!")
			fmt.Println("a: add an item to your to do list")
			fmt.Println("l: list all to do items")
			fmt.Println("t: toggle the status of a to do item")
			fmt.Println("d: delete an item from the to do list")
			fmt.Println("h: print this help prompt")
			fmt.Println("q: quit this program")
		case "q":
			return
		default:
			fmt.Printf("%s is not a valid command, type 'h' for a list of valid commands\n", command)
		}
	}
}

func main() {
	fmt.Println("Welcome to GO To Do!")
	database := Setup()
	promptAndRead(database)
}
