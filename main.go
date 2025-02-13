package main

import (
	"fmt"
	"os"

	"github.com/timmyha/todo-cli/cmd"
)

func main() {
	if len(os.Args) < 2 {
		cmd.OpenEditor("")
		return
	}

	switch os.Args[1] {
	case "add":
		if len(os.Args) < 4 {
			fmt.Println("Usage: todo add <day> \"task description\"")
			return
		}
		cmd.AddTask(os.Args[2], os.Args[3])
	case "complete":
		if len(os.Args) < 3 {
			fmt.Println("Usage: todo complete <day>")
			return
		}
		cmd.CompleteTask(os.Args[2])
	case "completed": // Ensure this case exists
		cmd.OpenEditor("completed")	
	default:
		fmt.Println("Unknown command. Use `add`, `complete`, or `completed`.")
	}
}

