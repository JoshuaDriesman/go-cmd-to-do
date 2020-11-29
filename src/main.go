package main

import "fmt"

func promptAndRead() {
	for {
		fmt.Print("Enter a command (type 'h' for a list of commands): ")

		var command string

		_, err := fmt.Scanln(&command)

		if err != nil {
			fmt.Println(err)
			return
		}

		switch command {
		case "h":
			fmt.Println("Here are all the commands this program supports!")
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
	promptAndRead()
}
