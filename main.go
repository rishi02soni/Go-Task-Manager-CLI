package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage:")
		fmt.Println(" add <task>")
		fmt.Println(" list")
		fmt.Println(" delete <id>")
		return
	}

	command := os.Args[1]

	switch command {

	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Please provide task title")
			return
		}
		addTask(os.Args[2])

	case "list":
		listTasks()

	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Provide task ID")
			return
		}
		id, _ := strconv.Atoi(os.Args[2])
		deleteTask(id)

	default:
		fmt.Println("Unknown command")
	}
}
