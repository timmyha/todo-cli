package cmd

import (
	"fmt"
	"github.com/timmyha/todo-cli/internal/storage"
)

func AddTask(day, task string) {
	err := storage.AppendTask(day, task)
	if err != nil {
		fmt.Println("Error adding task:", err)
	} else {
		fmt.Println("Task added to", day)
	}
}

